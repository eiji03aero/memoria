package service

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/model"
)

type Timeline struct {
	reg              interfaces.Registry
	timelinePostRepo repository.TimelinePost
	tptrRepo         repository.TimelinePostThreadRelation
	threadSvc        svc.Thread
}

func NewTimeline(reg interfaces.Registry) svc.Timeline {
	return &Timeline{
		reg:              reg,
		timelinePostRepo: reg.NewTimelinePostRepository(),
		tptrRepo:         reg.NewTimelinePostThreadRelationRepository(),
		threadSvc:        reg.NewThreadService(),
	}
}

func (s *Timeline) CreatePost(dto svc.TimelineCreatePostDTO) (tp *model.TimelinePost, err error) {
	id := GenerateUlid()

	tp, err = s.timelinePostRepo.Create(repository.TimelinePostCreateDTO{
		ID:          id,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	})
	if err != nil {
		return
	}

	thread, err := s.threadSvc.Create(svc.ThreadCreateDTO{
		UserSpaceID: dto.UserSpaceID,
	})
	if err != nil {
		return
	}

	_, err = s.tptrRepo.Create(repository.TimelinePostThreadRelationCreateDTO{
		TimelinePostID: id,
		ThreadID:       thread.ID,
	})
	if err != nil {
		return
	}

	err = s.threadSvc.AddPost(svc.ThreadAddPostDTO{
		ThreadID:    thread.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
		MediumIDs:   dto.MediumIDs,
	})
	if err != nil {
		return
	}

	return
}
