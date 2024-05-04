package service

import (
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
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

func (s *UserSpace) FindByID(id string) (ui *model.UserSpace, err error) {
	ui, err = s.userSpaceRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	})
	return
}
