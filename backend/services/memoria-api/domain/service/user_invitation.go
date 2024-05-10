package service

import (
	"memoria-api/domain/interfaces/repository"
)

type UserInvitation struct {
	userInvitationRepo repository.UserInvitation
}

type NewUserInvitationDTO struct {
	UserInvitationRepository repository.UserInvitation
}

func NewUserInvitation(dto NewUserInvitationDTO) *UserInvitation {
	return &UserInvitation{
		userInvitationRepo: dto.UserInvitationRepository,
	}
}
