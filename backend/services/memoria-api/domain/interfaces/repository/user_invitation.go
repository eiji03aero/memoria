package repository

import (
	"memoria-api/domain/model"
)

type UserInvitation interface {
	FindOne(findOption *FindOption) (ui *model.UserInvitation, err error)
	Create(dto UserInvitationCreateDTO) (err error)
}

type UserInvitationCreateDTO struct {
	ID     string
	UserID string
	Type   string
}
