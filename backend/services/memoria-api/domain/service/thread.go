package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/model"
)

type Thread struct {
	threadRepo   repository.Thread
	tmprRepo     repository.ThreadMicroPostRelation
	microPostSvc svc.MicroPost
}

func NewThread(reg interfaces.Registry) svc.Thread {
	return &Thread{
		threadRepo:   reg.NewThreadRepository(),
		tmprRepo:     reg.NewThreadMicroPostRelationRepository(),
		microPostSvc: reg.NewMicroPostService(),
	}
}

func (s *Thread) Create(dto svc.ThreadCreateDTO) (t *model.Thread, err error) {
	id := GenerateUlid()
	return s.threadRepo.Create(repository.ThreadCreateDTO{
		ID:          id,
		UserSpaceID: dto.UserSpaceID,
	})
}

func (s *Thread) AddPost(dto svc.ThreadAddPostDTO) (err error) {
	post, err := s.microPostSvc.Create(svc.MicroPostCreateDTO{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
		MediumIDs:   dto.MediumIDs,
	})
	if err != nil {
		return
	}

	_, err = s.tmprRepo.Create(repository.ThreadMicroPostRelationDTO{
		ThreadID:    dto.ThreadID,
		MicroPostID: post.ID,
	})
	if err != nil {
		return
	}

	return
}
