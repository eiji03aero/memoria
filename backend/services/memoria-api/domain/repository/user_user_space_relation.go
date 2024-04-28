package repository

type UserUserSpaceRelation interface {
	Create(dto UserUserSpaceRelationCreateDTO) (err error)
}

type UserUserSpaceRelationCreateDTO struct {
	UserID      string
	UserSpaceID string
}
