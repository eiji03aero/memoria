package repository

import "memoria-api/domain/model"

type MicroPost interface {
	Find(fOpt *FindOption) (mps []*model.MicroPost, err error)
	Create(dto MicroPostCreateDTO) (mp *model.MicroPost, err error)
}

type MicroPostCreateDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Content     string
}
