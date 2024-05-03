package repository

import (
	"memoria-api/domain/model"
)

type User interface {
	FindByID(dto UserFindByIDDTO) (user *model.User, err error)
	Create(dto UserCreateDTO) (err error)
	Update(user *model.User) (err error)
}

type UserFindByIDDTO struct {
	ID string
}

type UserCreateDTO struct {
	ID            string
	AccountStatus string
	Name          string
	Email         string
	PasswordHash  string
}
