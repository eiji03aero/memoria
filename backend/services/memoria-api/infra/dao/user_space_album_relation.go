package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type userSpaceAlbumRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUserSpaceAlbumRelation(db *gorm.DB) repository.UserSpaceAlbumRelation {
	return &userSpaceAlbumRelation[tbl.UserSpaceAlbumRelation]{db: db}
}

func (d *userSpaceAlbumRelation[T]) Find(findOption *repository.FindOption) (usars []*model.UserSpaceAlbumRelation, err error) {
	usarTbls := []*tbl.UserSpaceAlbumRelation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &usarTbls,
		name:       "user-space-album-relation",
	})
	if err != nil {
		return
	}

	usars = make([]*model.UserSpaceAlbumRelation, 0, len(usarTbls))
	for _, usarTbl := range usarTbls {
		usars = append(usars, usarTbl.ToModel())
	}
	return
}

func (d *userSpaceAlbumRelation[T]) FindOne(findOption *repository.FindOption) (usar *model.UserSpaceAlbumRelation, err error) {
	usarTbl := tbl.UserSpaceAlbumRelation{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &usarTbl,
		name:       "user-space-album-relation",
	})
	if err != nil {
		return
	}

	usar = usarTbl.ToModel()
	return
}

func (d *userSpaceAlbumRelation[T]) Create(dto repository.UserSpaceAlbumRelationCreateDTO) (usar *model.UserSpaceAlbumRelation, err error) {
	usarTbl := &tbl.UserSpaceAlbumRelation{
		UserSpaceID: dto.UserSpaceID,
		AlbumID:     dto.AlbumID,
	}

	err = d.db.Create(usarTbl).Error
	if err != nil {
		return
	}

	usar = usarTbl.ToModel()
	return
}
