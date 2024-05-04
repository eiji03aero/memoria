package service

import (
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
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

func (s *UserInvitation) FindByID(id string) (ui *model.UserInvitation, err error) {
	ui, err = s.userInvitationRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	})
	return
}
