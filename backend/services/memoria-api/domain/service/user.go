package service

import (
	"errors"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
)

type User struct {
	userRepo repository.User
}

type NewUserDTO struct {
	UserRepository repository.User
}

func NewUser(dto NewUserDTO) *User {
	return &User{
		userRepo: dto.UserRepository,
	}
}

func (s *User) FindByID(id string) (u *model.User, err error) {
	u, err = s.userRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	})
	return
}

func (s *User) FindByEmail(email string) (u *model.User, err error) {
	u, err = s.userRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "email = ?", Value: email},
		},
	})
	return
}

func (s *User) ExistsByEmail(email string) (exists bool, err error) {
	_, err = s.userRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "email = ?", Value: email},
		},
	})
	if errors.As(err, &cerrors.ResourceNotFound{}) {
		exists = false
		err = nil
		return
	}

	exists = true
	return
}

func (s *User) HasValidStatusForUse(userID string) (ok bool, err error) {
	user, err := s.FindByID(userID)
	if err != nil {
		return
	}

	ok = user.IsStatusValidForUse()
	return
}
