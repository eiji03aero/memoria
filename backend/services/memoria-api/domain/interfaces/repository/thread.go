package repository

import "memoria-api/domain/model"

type Thread interface {
	Find(fOpt *FindOption) (ts []*model.Thread, err error)
	FindOne(fOpt *FindOption) (t *model.Thread, err error)
	Create(dto ThreadCreateDTO) (t *model.Thread, err error)
}

type ThreadCreateDTO struct {
	ID          string
	UserSpaceID string
}
