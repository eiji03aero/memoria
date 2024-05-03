package repository

import (
	"memoria-api/domain/model"
)

type User interface {
	Find(dto UserFindDTO) (users []*model.User, err error)
	FindByID(dto UserFindByIDDTO) (user *model.User, err error)
	Create(dto UserCreateDTO) (err error)
	Update(user *model.User) (err error)
}

type UserFindDTO struct {
	FindOption *FindOption
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
