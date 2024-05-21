package model

import "time"

type UserSpaceAlbumRelation struct {
	UserSpaceID string
	AlbumID     string
}

type NewUserSpaceAlbumRelationDTO struct {
	UserSpaceID string
	AlbumID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUserSpaceAlbumRelation(dto NewUserSpaceAlbumRelationDTO) *UserSpaceAlbumRelation {
	return &UserSpaceAlbumRelation{
		UserSpaceID: dto.UserSpaceID,
		AlbumID:     dto.AlbumID,
	}
}
