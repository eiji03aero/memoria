package service

import (
	"memoria-api/domain/interfaces/repository"
)

type UserSpace struct {
	userSpaceRepo repository.UserSpace
}

type NewUserSpaceDTO struct {
	UserSpaceRepository repository.UserSpace
}

func NewUserSpace(dto NewUserSpaceDTO) *UserSpace {
	return &UserSpace{
		userSpaceRepo: dto.UserSpaceRepository,
	}
}
