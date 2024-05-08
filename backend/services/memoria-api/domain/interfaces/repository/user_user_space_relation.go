package repository

import "memoria-api/domain/model"

type UserUserSpaceRelation interface {
	Find(findOption *FindOption) ([]*model.UserUserSpaceRelation, error)
	FindOne(findOption *FindOption) (*model.UserUserSpaceRelation, error)
	Create(dto UserUserSpaceRelationCreateDTO) (err error)
}

type UserUserSpaceRelationCreateDTO struct {
	UserID      string
	UserSpaceID string
}
