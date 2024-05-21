package repository

import "memoria-api/domain/model"

type AlbumMediumRelation interface {
	Find(findOption *FindOption) ([]*model.AlbumMediumRelation, error)
	FindOne(findOption *FindOption) (*model.AlbumMediumRelation, error)
	Exists(findOption *FindOption) (bool, error)
	Create(dto AlbumMediumRelationCreateDTO) (err error)
	Delete(dto AlbumMediumRelationDeleteDTO) (err error)
}

type AlbumMediumRelationCreateDTO struct {
	AlbumID  string
	MediumID string
}

type AlbumMediumRelationDeleteDTO struct {
	AlbumID  string
	MediumID string
}
