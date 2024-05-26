package usecase

import (
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/model"
	"memoria-api/domain/service"
)

type Album interface {
	Create(dto AlbumCreateDTO) (ret AlbumCreateRet, err error)
	Delete(dto AlbumDeleteDTO) (err error)
	AddMedia(dto AlbumAddMediaDTO) (err error)
	RemoveMedia(dto AlbumRemoveMediaDTO) (err error)
}

type album struct {
	reg        interfaces.Registry
	albumRepo  repository.Album
	usarRepo   repository.UserSpaceAlbumRelation
	mediumRepo repository.Medium
	amrRepo    repository.AlbumMediumRelation
	albumSvc   svc.Album
}

func NewAlbum(reg interfaces.Registry) (u Album) {
	u = &album{
		reg:        reg,
		albumRepo:  reg.NewAlbumRepository(),
		usarRepo:   reg.NewUserSpaceAlbumRelationRepository(),
		mediumRepo: reg.NewMediumRepository(),
		amrRepo:    reg.NewAlbumMediumRelationRepository(),
		albumSvc:   reg.NewAlbumService(),
	}
	return
}

type AlbumCreateDTO struct {
	Name        *string
	UserSpaceID string
}

type AlbumCreateRet struct {
	Album *model.Album
}

func (u *album) Create(dto AlbumCreateDTO) (ret AlbumCreateRet, err error) {
	// -------------------- validation --------------------
	if dto.Name == nil {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Required,
			Name: "name",
		})
		return
	}

	// -------------------- execution --------------------
	albumID := service.GenerateUlid()
	album, err := u.albumRepo.Create(repository.AlbumCreateDTO{
		ID:   albumID,
		Name: *dto.Name,
	})
	if err != nil {
		return
	}

	_, err = u.usarRepo.Create(repository.UserSpaceAlbumRelationCreateDTO{
		UserSpaceID: dto.UserSpaceID,
		AlbumID:     albumID,
	})
	if err != nil {
		return
	}

	ret.Album = album
	return
}

type AlbumDeleteDTO struct {
	ID string
}

func (u *album) Delete(dto AlbumDeleteDTO) (err error) {
	mediumUc := NewMedium(u.reg)

	amrs, err := u.amrRepo.Find(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "album_id = ?", Value: dto.ID},
		},
	})
	if err != nil {
		return
	}

	// delete amr
	err = u.amrRepo.Delete(repository.AlbumMediumRelationDeleteDTO{
		AlbumID: dto.ID,
	})
	if err != nil {
		return
	}

	// delete media
	for _, amr := range amrs {
		e := mediumUc.Delete(MediumDeleteDTO{
			MediumID: amr.MediumID,
		})
		if e != nil {
			err = e
			return
		}
	}

	// delete album
	err = u.albumRepo.Delete(repository.AlbumDeleteDTO{
		ID: dto.ID,
	})

	return
}

type AlbumAddMediaDTO struct {
	UserSpaceID string
	AlbumIDs    []string
	MediumIDs   []string
}

func (u *album) AddMedia(dto AlbumAddMediaDTO) (err error) {
	for _, albumID := range dto.AlbumIDs {
		e := u.albumSvc.AddMedia(svc.AlbumAddMediaDTO{
			UserSpaceID: dto.UserSpaceID,
			AlbumID:     albumID,
			MediumIDs:   dto.MediumIDs,
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}

type AlbumRemoveMediaDTO struct {
	AlbumIDs  []string
	MediumIDs []string
}

func (u *album) RemoveMedia(dto AlbumRemoveMediaDTO) (err error) {
	for _, albumID := range dto.AlbumIDs {
		e := u.albumSvc.RemoveMedia(svc.AlbumRemoveMediaDTO{
			AlbumID:   albumID,
			MediumIDs: dto.MediumIDs,
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}
