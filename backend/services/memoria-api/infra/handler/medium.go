package handler

import (
	"net/http"

	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/handler/res"

	"github.com/gin-gonic/gin"
)

type Medium struct{}

func NewMedium() *Medium {
	return &Medium{}
}

type MediumFindReq struct {
	AlbumID *string `form:"album_id"`
}

type MediumFindRes struct {
	Media []*res.Medium `json:"media"`
}

func (h *Medium) Find(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	mediumRepo := reg.NewMediumRepository()

	query := MediumFindReq{}
	err = c.ShouldBindQuery(&query)
	if err != nil {
		return
	}

	findOpt := &repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "user_space_id = ?", Value: cctx.GetUserSpaceID()},
		},
		Joins: []*repository.FindOptionJoin{
			{Query: "left join album_medium_relations amr on amr.medium_id = media.id"},
		},
		Order: "media.created_at desc",
	}
	if query.AlbumID != nil {
		findOpt.Filters = append(findOpt.Filters, &repository.FindOptionFilter{
			Query: "amr.album_id = ?", Value: *query.AlbumID,
		})
	}

	media, err := mediumRepo.Find(findOpt)
	if err != nil {
		return
	}

	mediumRes := MediumFindRes{}
	mediumRes.Media = make([]*res.Medium, 0, len(media))
	for _, medium := range media {
		mediumRes.Media = append(mediumRes.Media, res.MediumFromModel(medium))
	}

	status = http.StatusOK
	data = mediumRes
	return
}

type MediumRequestUploadURLsReq struct {
	FileNames []string `json:"file_names"`
	AlbumIDs  []string `json:"album_ids"`
}

type MediumRequestUploadURLsRes struct {
	UploadURLs []res.UploadURL `json:"upload_urls"`
}

func (h *Medium) RequestUploadURLs(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	mediumUc := usecase.NewMedium(reg)

	body := MediumRequestUploadURLsReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	urlsRet, err := mediumUc.RequestUploadURLs(usecase.MediumRequestUploadURLsDTO{
		UserID:      cctx.GetUserID(),
		UserSpaceID: cctx.GetUserSpaceID(),
		FileNames:   body.FileNames,
		AlbumIDs:    body.AlbumIDs,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	mediumRes := MediumRequestUploadURLsRes{}
	for _, purl := range urlsRet.PresignedURLs {
		mediumRes.UploadURLs = append(mediumRes.UploadURLs, res.UploadURL{URL: purl.URL, MediumID: purl.MediumID})
	}

	data = mediumRes
	return
}

type MediumConfirmUploadsReq struct {
	MediumIDs []string `json:"medium_ids"`
}

func (h *Medium) ConfirmUploads(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	// TODO utilize this?
	// cctx := ccontext.NewContext(c)
	mediumUc := usecase.NewMedium(reg)

	body := MediumConfirmUploadsReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	// TODO: turn this into bgjob
	err = mediumUc.ConfirmUploads(usecase.MediumConfirmUploadsDTO{
		MediumIDs: body.MediumIDs,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	return
}

func (h *Medium) Delete(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	mediumUc := usecase.NewMedium(reg)

	mediumID := c.Param("id")

	// TODO: turn this into bgjob
	err = mediumUc.Delete(usecase.MediumDeleteDTO{
		MediumID: mediumID,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	return
}
