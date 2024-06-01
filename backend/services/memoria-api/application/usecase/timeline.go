package usecase

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/svc"
)

type Timeline interface {
	CreatePost(dto TimelineCreatePostDTO) (ret TimelineCreatePostRet, err error)
}

type timeline struct {
	reg         interfaces.Registry
	timelineSvc svc.Timeline
}

func NewTimeline(reg interfaces.Registry) (u Timeline) {
	u = &timeline{
		reg:         reg,
		timelineSvc: reg.NewTimelineService(),
	}
	return
}

type TimelineCreatePostDTO struct {
	UserID      string
	UserSpaceID string
	Content     string
	MediumIDs   []string
}

type TimelineCreatePostRet struct{}

func (u *timeline) CreatePost(dto TimelineCreatePostDTO) (ret TimelineCreatePostRet, err error) {
	_, err = u.timelineSvc.CreatePost(svc.TimelineCreatePostDTO{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
		MediumIDs:   dto.MediumIDs,
	})
	return
}
