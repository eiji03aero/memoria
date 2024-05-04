package repository

import "memoria-api/domain/model"

type UserSpace interface {
	FindOne(findOption *FindOption) (userSpace *model.UserSpace, err error)
	Create(dto UserSpaceCreateDTO) (err error)
}

type UserSpaceCreateDTO struct {
	ID   string
	Name string
}
