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

type UserIsEmailTakenDTO struct {
	Email string
}

func (s *User) IsEmailTaken(dto UserIsEmailTakenDTO) (taken bool, err error) {
	users, err := s.userRepo.Find(repository.UserFindDTO{
		FindOption: &repository.FindOption{
			Filters: []*repository.FindOptionFilter{
				{Query: "email = ?", Value: dto.Email},
			},
		},
	})
	if err != nil {
		return
	}

	taken = len(users) > 0
	return
}
