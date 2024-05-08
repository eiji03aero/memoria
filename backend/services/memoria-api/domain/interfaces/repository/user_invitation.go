package repository

import (
	"memoria-api/domain/model"
)

type UserInvitation interface {
	Find(findOption *FindOption) (uis []*model.UserInvitation, err error)
	FindOne(findOption *FindOption) (ui *model.UserInvitation, err error)
	Create(dto UserInvitationCreateDTO) (err error)
}

type UserInvitationCreateDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Type        string
}
