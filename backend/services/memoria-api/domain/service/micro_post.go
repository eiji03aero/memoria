package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/model"
)

type MicroPost struct {
	microPostRepo repository.MicroPost
	mpmrRepo      repository.MicroPostMediumRelation
}

func NewMicroPost(reg interfaces.Registry) svc.MicroPost {
	return &MicroPost{
		microPostRepo: reg.NewMicroPostRepository(),
		mpmrRepo:      reg.NewMicroPostMediumRelationRepository(),
	}
}

func (s *MicroPost) Create(dto svc.MicroPostCreateDTO) (mp *model.MicroPost, err error) {
	id := GenerateUlid()

	mp, err = s.microPostRepo.Create(repository.MicroPostCreateDTO{
		ID:          id,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
	})
	if err != nil {
		return
	}

	for _, mediumID := range dto.MediumIDs {
		_, e := s.mpmrRepo.Create(repository.MicroPostMediumRelationCreateDTO{
			MicroPostID: id,
			MediumID:    mediumID,
		})
		if e != nil {
			err = e
			return
		}
	}

	return
}
