package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
)

type Album struct {
	mediumRepo repository.Medium
	amrRepo    repository.AlbumMediumRelation
}

func NewAlbum(reg interfaces.Registry) svc.Album {
	return &Album{
		mediumRepo: reg.NewMediumRepository(),
		amrRepo:    reg.NewAlbumMediumRelationRepository(),
	}
}

func (s *Album) AddMedia(dto svc.AlbumAddMediaDTO) (err error) {
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
