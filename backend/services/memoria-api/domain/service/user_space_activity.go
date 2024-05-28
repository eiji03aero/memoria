package service

import (
	"encoding/json"

	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/model"
	"memoria-api/domain/value"
)

type UserSpaceActivity struct {
	usaRepo repository.UserSpaceActivity
}

func NewUserSpaceActivity(reg interfaces.Registry) svc.UserSpaceActivity {
	return &UserSpaceActivity{
		usaRepo: reg.NewUserSpaceActivityRepository(),
	}
}

func (s *UserSpaceActivity) CreateInviteUserJoined(dto svc.UserSpaceActivityCreateInviteUserJoined) (err error) {
	id := GenerateUlid()

	bs, err := json.Marshal(model.UserSpaceActivityData_InviteUserJoined{
		UserID: dto.UserID,
	})
	if err != nil {
		return
	}

	_, err = s.usaRepo.Create(repository.UserSpaceActivityCreateDTO{
		ID:          id,
		Type:        value.UserSpaceActivityType_InviteUserJoined,
		UserSpaceID: dto.UserSpaceID,
		Data:        string(bs),
	})
	return
}

func (s *UserSpaceActivity) CreateUserUploadedMedia(dto svc.UserSpaceActivityCreateUserUploadedMedia) (err error) {
	id := GenerateUlid()

	bs, err := json.Marshal(model.UserSpaceActivityData_UserUploadedMedia{
		UserID:    dto.UserID,
		MediumIDs: dto.MediumIDs,
	})
	if err != nil {
		return
	}

	_, err = s.usaRepo.Create(repository.UserSpaceActivityCreateDTO{
		ID:          id,
		Type:        value.UserSpaceActivityType_UserUploadedMedia,
		UserSpaceID: dto.UserSpaceID,
		Data:        string(bs),
	})
	return
}
