package service

import (
	"memoria-api/domain/interfaces/repository"
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
