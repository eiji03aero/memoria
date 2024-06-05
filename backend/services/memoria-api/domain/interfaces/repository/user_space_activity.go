package repository

import (
	"memoria-api/domain/model"
	"memoria-api/domain/value"
)

type UserSpaceActivity interface {
	Find(findOpt *FindOption) (usas []*model.UserSpaceActivity, err error)
	FindOne(findOpt *FindOption) (usa *model.UserSpaceActivity, err error)
	FindOneByID(id string) (usa *model.UserSpaceActivity, err error)
	Create(dto UserSpaceActivityCreateDTO) (usa *model.UserSpaceActivity, err error)
}

type UserSpaceActivityCreateDTO struct {
	ID          string
	UserSpaceID string
	Type        value.UserSpaceActivityType
	Data        string
}
