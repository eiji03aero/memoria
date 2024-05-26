package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type album[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewAlbum(db *gorm.DB) repository.Album {
	return &album[tbl.Album]{db: db}
}

func (d *album[T]) Find(findOption *repository.FindOption) (albums []*model.Album, err error) {
	albumTbls := []*tbl.Album{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &albumTbls,
		name:       "album",
	})
	if err != nil {
		return
	}

	albums = make([]*model.Album, 0, len(albumTbls))
	for _, albumTbl := range albumTbls {
		albums = append(albums, albumTbl.ToModel())
	}
	return
}

func (d *album[T]) FindOne(findOption *repository.FindOption) (album *model.Album, err error) {
	albumTbl := tbl.Album{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &albumTbl,
		name:       "album",
	})
	if err != nil {
		return
	}

	album = albumTbl.ToModel()
	return
}

func (d *album[T]) FindOneByID(id string) (album *model.Album, err error) {
	return d.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	})
}

func (d *album[T]) Create(dto repository.AlbumCreateDTO) (album *model.Album, err error) {
	albumTbl := &tbl.Album{
		ID:   dto.ID,
		Name: dto.Name,
	}

	err = d.db.Create(albumTbl).Error
	if err != nil {
		return
	}

	album = albumTbl.ToModel()
	return
}

func (d *album[T]) Delete(dto repository.AlbumDeleteDTO) (err error) {
	albumTbl := &tbl.Album{
		ID: dto.ID,
	}

	err = d.db.Delete(albumTbl).Error
	if err != nil {
		return
	}

	return
}
