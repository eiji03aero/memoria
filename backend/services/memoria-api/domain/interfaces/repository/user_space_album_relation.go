package repository

import "memoria-api/domain/model"

type UserSpaceAlbumRelation interface {
	Find(findOpt *FindOption) ([]*model.UserSpaceAlbumRelation, error)
	FindOne(findOpt *FindOption) (*model.UserSpaceAlbumRelation, error)
	Create(dto UserSpaceAlbumRelationCreateDTO) (usar *model.UserSpaceAlbumRelation, err error)
}

type UserSpaceAlbumRelationCreateDTO struct {
	UserSpaceID string
	AlbumID     string
}
