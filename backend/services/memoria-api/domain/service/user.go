package service

import (
	"memoria-api/domain/interfaces/repository"
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

func (s *User) HasValidStatusForUse(userID string) (ok bool, err error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return
	}

	ok = user.IsStatusValidForUse()
	return
}
