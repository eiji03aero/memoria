package model

import "time"

type UserUserSpaceRelation struct {
	UserID      string
	UserSpaceID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewUserUserSpaceRelationDTO struct {
	UserID      string
	UserSpaceID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUserUserSpaceRelation(dto NewUserUserSpaceRelationDTO) *UserUserSpaceRelation {
	return &UserUserSpaceRelation{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	}
}
