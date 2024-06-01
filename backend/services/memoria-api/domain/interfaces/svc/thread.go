package svc

import "memoria-api/domain/model"

type Thread interface {
	Create(dto ThreadCreateDTO) (t *model.Thread, err error)
	AddPost(dto ThreadAddPostDTO) (err error)
}

type ThreadCreateDTO struct {
	UserSpaceID string
}

type ThreadAddPostDTO struct {
	ThreadID    string
	UserID      string
	UserSpaceID string
	Content     string
	MediumIDs   []string
}
