package repository

import "memoria-api/domain/model"

type Album interface {
	Find(findOpt *FindOption) (albums []*model.Album, err error)
	FindOne(findOpt *FindOption) (album *model.Album, err error)
	FindOneByID(id string) (album *model.Album, err error)
	Create(dto AlbumCreateDTO) (album *model.Album, err error)
	Delete(dto AlbumDeleteDTO) (err error)
}

type AlbumCreateDTO struct {
	ID   string
	Name string
}

type AlbumDeleteDTO struct {
	ID string
}
