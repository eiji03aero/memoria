package repository

import "memoria-api/domain/model"

type UserSpace interface {
	FindByID(dto UserSpaceFindByIDDTO) (userSpace *model.UserSpace, err error)
	Create(dto UserSpaceCreateDTO) (err error)
}

type UserSpaceFindByIDDTO struct {
	ID string
}

type UserSpaceCreateDTO struct {
	ID   string
	Name string
}
