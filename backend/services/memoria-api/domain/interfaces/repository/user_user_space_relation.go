package repository

import "memoria-api/domain/model"

type UserUserSpaceRelation interface {
	FindOne(findOption *FindOption) (*model.UserUserSpaceRelation, error)
	Create(dto UserUserSpaceRelationCreateDTO) (err error)
}

type UserUserSpaceRelationCreateDTO struct {
	UserID      string
	UserSpaceID string
}
