package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
)

type Medium struct {
	s3Client   interfaces.S3Client
	mediumRepo repository.Medium
}

type NewMediumDTO struct {
	S3Client   interfaces.S3Client
	MediumRepo repository.Medium
}

func NewMedium(dto NewMediumDTO) svc.Medium {
	return &Medium{
		s3Client:   dto.S3Client,
		mediumRepo: dto.MediumRepo,
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

	err = s.s3Client.DeleteFolder(interfaces.S3ClientDeleteFolderDTO{
		Prefix: medium.GetFileDirectoryPath(),
	})

	err = s.mediumRepo.Delete(repository.MediumDeleteDTO{
		ID: dto.MediumID,
	})
	return
}
