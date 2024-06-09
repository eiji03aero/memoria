package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
)

type Medium struct {
	s3Client   interfaces.S3Client
	mediumRepo repository.Medium
	amrRepo    repository.AlbumMediumRelation
}

func NewMedium(reg interfaces.Registry) svc.Medium {
	return &Medium{
		s3Client:   reg.NewS3Client(),
		mediumRepo: reg.NewMediumRepository(),
		amrRepo:    reg.NewAlbumMediumRelationRepository(),
	}
}

func (s *Medium) Delete(dto svc.MediumDeleteDTO) (err error) {
	medium, err := s.mediumRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: dto.MediumID},
		},
	})
	if err != nil {
		return
	}

	err = s.amrRepo.Delete(repository.AlbumMediumRelationDeleteDTO{
		MediumID: dto.MediumID,
	})
	if err != nil {
		return
	}

	err = s.s3Client.DeleteFolder(interfaces.S3ClientDeleteFolderDTO{
		Prefix: medium.GetFileDirectoryPath(),
	})

	err = s.mediumRepo.Delete(repository.MediumDeleteDTO{
		ID: dto.MediumID,
	})
	return
}
