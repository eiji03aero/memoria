package model

import "memoria-api/domain/value"

type UserInvitation struct {
	ID     string
	UserID string
	Type   value.UserInvitationType
}

type NewUserInvitationDTO struct {
	ID     string
	UserID string
	Type   value.UserInvitationType
}

func NewUserInvitation(dto NewUserInvitationDTO) (*UserInvitation, error) {
	return &UserInvitation{
		ID:     dto.ID,
		UserID: dto.UserID,
		Type:   dto.Type,
	}, nil
}
