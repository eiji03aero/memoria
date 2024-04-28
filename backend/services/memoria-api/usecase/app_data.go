package usecase

import (
	"memoria-api/domain/model"
	"memoria-api/domain/repository"
	"memoria-api/registry"
)

type AppData interface {
	Get(dto AppDataGetDTO) (AppDataGetRet, error)
}

type appData struct {
	userRepo      repository.User
	userSpaceRepo repository.UserSpace
}

func NewAppData(reg registry.Registry) AppData {
	return &appData{
		userRepo:      reg.NewUserRepository(),
		userSpaceRepo: reg.NewUserSpaceRepository(),
	}
}

type AppDataGetDTO struct {
	UserID      string
	UserSpaceID string
}
type AppDataGetRet struct {
	User      *model.User
	UserSpace *model.UserSpace
}

func (u *appData) Get(dto AppDataGetDTO) (ret AppDataGetRet, err error) {
	user, err := u.userRepo.FindByID(repository.UserFindByIDDTO{
		ID: dto.UserID,
	})
	if err != nil {
		return
	}

	userSpace, err := u.userSpaceRepo.FindByID(repository.UserSpaceFindByIDDTO{
		ID: dto.UserSpaceID,
	})
	if err != nil {
		return
	}

	ret = AppDataGetRet{
		User:      user,
		UserSpace: userSpace,
	}

	return
}
