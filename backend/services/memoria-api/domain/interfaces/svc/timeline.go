package svc

import "memoria-api/domain/model"

type Timeline interface {
	CreatePost(dto TimelineCreatePostDTO) (tp *model.TimelinePost, err error)
}

type TimelineCreatePostDTO struct {
	UserID      string
	UserSpaceID string
	Content     string
	MediumIDs   []string
}
