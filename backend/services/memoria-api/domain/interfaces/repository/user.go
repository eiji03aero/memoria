package repository

import (
	"memoria-api/domain/model"
)

type User interface {
	Find(findOpt *FindOption) (users []*model.User, err error)
	FindOne(findOpt *FindOption) (user *model.User, err error)
	Create(dto UserCreateDTO) (err error)
	Update(user *model.User) (err error)
}

type UserCreateDTO struct {
	ID            string
	AccountStatus string
	Name          string
	Email         string
	PasswordHash  string
}
