package service

import (
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
)

type UserUserSpaceRelation struct {
	userUserSpaceRelationRepo repository.UserUserSpaceRelation
}

type NewUserUserSpaceRelationDTO struct {
	UserUserSpaceRelationRepository repository.UserUserSpaceRelation
}

func NewUserUserSpaceRelation(dto NewUserUserSpaceRelationDTO) *UserUserSpaceRelation {
	return &UserUserSpaceRelation{
		userUserSpaceRelationRepo: dto.UserUserSpaceRelationRepository,
	}
}

func (s *UserUserSpaceRelation) FindByUserID(userID string) (uusr *model.UserUserSpaceRelation, err error) {
	return s.userUserSpaceRelationRepo.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "user_id = ?", Value: userID},
		},
	})
}
