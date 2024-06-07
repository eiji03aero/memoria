package usecase

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/service"
	"memoria-api/util"

	"github.com/davidbyttow/govips/v2/vips"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type Medium interface {
	RequestUploadURLs(dto MediumRequestUploadURLsDTO) (MediumRequestUploadURLsRet, error)
	ConfirmUploads(dto MediumConfirmUploadsDTO) error
	CreateThumbnails(dto MediumCreateThumbnailsDTO) error
	Delete(dto MediumDeleteDTO) error
}

type medium struct {
	bgjobInvoker            interfaces.BGJobInvoker
	s3Client                interfaces.S3Client
	mediumRepo              repository.Medium
	albumMediumRelationRepo repository.AlbumMediumRelation
	mediumSvc               svc.Medium
	usaSvc                  svc.UserSpaceActivity
}

func NewMedium(reg interfaces.Registry) (u Medium) {
	return &medium{
		bgjobInvoker:            reg.NewBGJobInvoker(),
		s3Client:                reg.NewS3Client(),
		mediumRepo:              reg.NewMediumRepository(),
		albumMediumRelationRepo: reg.NewAlbumMediumRelationRepository(),
		mediumSvc:               reg.NewMediumService(),
		usaSvc:                  reg.NewUserSpaceActivityService(),
	}
}

type MediumRequestUploadURLsDTO struct {
	UserID      string
	UserSpaceID string
	FileNames   []string
	AlbumIDs    []string
}

type MediumUploadURL struct {
	URL      string
	MediumID string
}

type MediumRequestUploadURLsRet struct {
	PresignedURLs []MediumUploadURL
}

func (u medium) RequestUploadURLs(dto MediumRequestUploadURLsDTO) (ret MediumRequestUploadURLsRet, err error) {
	ret.PresignedURLs = make([]MediumUploadURL, 0, len(dto.FileNames))
	for _, fileName := range dto.FileNames {
		mediumID := service.GenerateUlid()
		ext := filepath.Ext(fileName)
		name := strings.TrimSuffix(fileName, ext)

		_, e := u.mediumRepo.Create(repository.MediumCreateDTO{
			ID:          mediumID,
			UserID:      dto.UserID,
			UserSpaceID: dto.UserSpaceID,
			Name:        name,
			Extension:   ext,
		})
		if e != nil {
			err = e
			return
		}

		for _, albumID := range dto.AlbumIDs {
			e = u.albumMediumRelationRepo.Create(repository.AlbumMediumRelationCreateDTO{
				AlbumID:  albumID,
				MediumID: mediumID,
			})
			if e != nil {
				err = e
				return
			}
		}

		presignedURL, e := u.s3Client.GetPresignedPutObjectURL(interfaces.S3ClientGetPresignedPutObjectURLDTO{
			Key: "media/" + mediumID + "/original" + filepath.Ext(fileName),
		})
		if e != nil {
			err = e
			return
		}

		ret.PresignedURLs = append(ret.PresignedURLs, MediumUploadURL{
			URL:      presignedURL,
			MediumID: mediumID,
		})
	}

	return
}

type MediumConfirmUploadsDTO struct {
	UserID      string
	UserSpaceID string
	MediumIDs   []string
}

func (u *medium) ConfirmUploads(dto MediumConfirmUploadsDTO) (err error) {
	for _, mediumID := range dto.MediumIDs {
		u.bgjobInvoker.CreateThumbnails(interfaces.BGJobInvokerCreateThumbnailsDTO{
			MediumID: mediumID,
		})
	}

	err = u.usaSvc.CreateUserUploadedMedia(svc.UserSpaceActivityCreateUserUploadedMedia{
		UserSpaceID: dto.UserSpaceID,
		UserID:      dto.UserID,
		MediumIDs:   dto.MediumIDs,
	})

	return
}

type MediumCreateThumbnailsDTO struct {
	MediumID string
}

func (u *medium) CreateThumbnails(dto MediumCreateThumbnailsDTO) (err error) {
	medium, err := u.mediumRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: dto.MediumID},
		},
	})
	if err != nil {
		return
	}

	// -------------------- download the original file --------------------
	mediumFile, err := util.DownloadFile(util.DownloadFileDTO{
		FileName: medium.Name + medium.Extension,
		URL:      medium.GetOriginalURL(),
	})
	defer mediumFile.Close()

	// -------------------- thumbnail creation for image --------------------
	if medium.IsImage() {
		e := u.CreateThumbnail(MediumCreateThumbnailDTO{
			MediumID:      medium.ID,
			Width:         240,
			ImageFilePath: mediumFile.Name(),
		})
		if e != nil {
			err = e
			return
		}
	} else if medium.IsVideo() {
		// -------------------- thumbnail creation for video --------------------
		buf := bytes.NewBuffer(nil)
		e := ffmpeg_go.Input(mediumFile.Name()).
			Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", 5)}).
			Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
			WithOutput(buf, os.Stdout).
			Run()
		if e != nil {
			err = e
			return e
		}

		thumbnail, e := os.CreateTemp("", "*-"+medium.ID+"-thumbnail.jpeg")
		if e != nil {
			err = e
			return
		}

		_, e = io.Copy(thumbnail, buf)
		if e != nil {
			err = e
			return
		}

		e = u.CreateThumbnail(MediumCreateThumbnailDTO{
			MediumID:      medium.ID,
			Width:         240,
			ImageFilePath: thumbnail.Name(),
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}

type MediumCreateThumbnailDTO struct {
	MediumID      string
	Width         int
	ImageFilePath string
}

func (u *medium) CreateThumbnail(dto MediumCreateThumbnailDTO) (err error) {
	vipsImg, e := vips.NewImageFromFile(dto.ImageFilePath)
	if e != nil {
		err = e
		return
	}

	vipsImg.Resize(
		float64(float64(dto.Width)/float64(vipsImg.Width())),
		vips.KernelAuto,
	)

	imgBytes, _, e := vipsImg.ExportPng(vips.NewPngExportParams())
	if e != nil {
		err = e
		return
	}

	err = u.s3Client.PutObject(interfaces.S3ClientPutObjectDTO{
		Key:  "media/" + dto.MediumID + "/" + strconv.Itoa(dto.Width) + ".png",
		Body: bytes.NewReader(imgBytes),
	})
	return
}

type MediumDeleteDTO struct {
	MediumID string
}

func (u *medium) Delete(dto MediumDeleteDTO) (err error) {
	err = u.albumMediumRelationRepo.Delete(repository.AlbumMediumRelationDeleteDTO{
		MediumID: dto.MediumID,
	})
	if err != nil {
		return
	}

	err = u.mediumSvc.Delete(svc.MediumDeleteDTO{
		MediumID: dto.MediumID,
	})
	return
}
