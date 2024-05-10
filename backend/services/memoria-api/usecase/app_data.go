package usecase

import (
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/domain/service"
	"memoria-api/registry"
)

type AppData interface {
	Get(dto AppDataGetDTO) (AppDataGetRet, error)
}

type appData struct {
	userRepo      repository.User
	userSpaceRepo repository.UserSpace
	userSvc       *service.User
	userSpaceSvc  *service.UserSpace
}

func NewAppData(reg registry.Registry) AppData {
	return &appData{
		userRepo:      reg.NewUserRepository(),
		userSpaceRepo: reg.NewUserSpaceRepository(),
		userSvc:       reg.NewUserService(),
		userSpaceSvc:  reg.NewUserSpaceService(),
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
	user, err := u.userRepo.FindByID(dto.UserID)
	if err != nil {
		return
	}

	userSpace, err := u.userSpaceRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: dto.UserSpaceID},
		},
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
