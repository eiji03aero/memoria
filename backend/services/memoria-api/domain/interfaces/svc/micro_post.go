package svc

import "memoria-api/domain/model"

type MicroPost interface {
	Create(dto MicroPostCreateDTO) (t *model.MicroPost, err error)
}

type MicroPostCreateDTO struct {
	UserID      string
	UserSpaceID string
	Content     string
	MediumIDs   []string
}
