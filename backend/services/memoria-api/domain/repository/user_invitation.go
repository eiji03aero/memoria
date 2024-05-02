package repository

import (
	"memoria-api/domain/model"
)

type UserInvitation interface {
	FindByID(dto UserInvitationFindByIDDTO) (ui *model.UserInvitation, err error)
	Create(dto UserInvitationCreateDTO) (err error)
}

type UserInvitationFindByIDDTO struct {
	ID string
}

type UserInvitationCreateDTO struct {
	ID     string
	UserID string
	Type   string
}
