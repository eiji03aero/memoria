package service

import (
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
)

type Album struct {
	mediumRepo repository.Medium
	albumRepo  repository.Album
	amrRepo    repository.AlbumMediumRelation
	usaRepo    repository.UserSpaceAlbumRelation
}

func NewAlbum(reg interfaces.Registry) svc.Album {
	return &Album{
		mediumRepo: reg.NewMediumRepository(),
		albumRepo:  reg.NewAlbumRepository(),
		amrRepo:    reg.NewAlbumMediumRelationRepository(),
		usaRepo:    reg.NewUserSpaceAlbumRelationRepository(),
	}
}

func (s *Album) AddMedia(dto svc.AlbumAddMediaDTO) (err error) {
	usa, err := s.usaRepo.FindOneByAlbumID(dto.AlbumID)
	if err != nil {
		return
	}
	if usa.UserSpaceID != dto.UserSpaceID {
		err = cerrors.NewValidation(cerrors.NewValidationDTO{
			Key:  cerrors.ValidationKey_Consistency,
			Name: "user-space-id",
		})
		return
	}

	for _, mediumID := range dto.MediumIDs {
		exists, e := s.amrRepo.Exists(&repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "album_id = ?", Value: dto.AlbumID},
				{Query: "medium_id = ?", Value: mediumID},
			},
		})
		if err != nil {
			err = e
			return
		}
		if exists {
			continue
		}

		e = s.amrRepo.Create(repository.AlbumMediumRelationCreateDTO{
			AlbumID:  dto.AlbumID,
			MediumID: mediumID,
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}

func (s *Album) RemoveMedia(dto svc.AlbumRemoveMediaDTO) (err error) {
	for _, mediumID := range dto.MediumIDs {
		e := s.amrRepo.Delete(repository.AlbumMediumRelationDeleteDTO{
			AlbumID:  dto.AlbumID,
			MediumID: mediumID,
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}
