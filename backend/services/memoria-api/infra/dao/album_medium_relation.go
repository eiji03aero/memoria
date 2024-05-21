package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type albumMediumRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewAlbumMediumRelation(db *gorm.DB) repository.AlbumMediumRelation {
	return &albumMediumRelation[tbl.AlbumMediumRelation]{db: db}
}

func (d *albumMediumRelation[T]) Find(findOption *repository.FindOption) (amrs []*model.AlbumMediumRelation, err error) {
	amrTbls := []*tbl.AlbumMediumRelation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &amrTbls,
		name:       "album-medium-relation",
	})
	if err != nil {
		return
	}

	amrs = make([]*model.AlbumMediumRelation, 0, len(amrTbls))
	for _, amrTbl := range amrTbls {
		amrs = append(amrs, amrTbl.ToModel())
	}
	return
}

func (d *albumMediumRelation[T]) FindOne(findOption *repository.FindOption) (amr *model.AlbumMediumRelation, err error) {
	amrTbl := tbl.AlbumMediumRelation{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &amrTbl,
		name:       "album-medium-relation",
	})
	if err != nil {
		return
	}

	amr = amrTbl.ToModel()
	return
}

func (d *albumMediumRelation[T]) Exists(findOpt *repository.FindOption) (exists bool, err error) {
	return d.exists(existsDTO{
		db:         d.db,
		findOption: findOpt,
		data:       &tbl.AlbumMediumRelation{},
		name:       "album-medium-relation",
	})
}

func (d *albumMediumRelation[T]) Create(dto repository.AlbumMediumRelationCreateDTO) (err error) {
	amr := &tbl.AlbumMediumRelation{
		AlbumID:  dto.AlbumID,
		MediumID: dto.MediumID,
	}

	result := d.db.Create(amr)
	err = result.Error
	return
}

func (d *albumMediumRelation[T]) Delete(dto repository.AlbumMediumRelationDeleteDTO) (err error) {
	amr := &tbl.AlbumMediumRelation{}
	query := d.db

	if dto.AlbumID != "" {
		query = query.Where("album_id = ?", dto.AlbumID)
	}
	if dto.MediumID != "" {
		query = query.Where("medium_id = ?", dto.MediumID)
	}

	err = query.Delete(amr).Error
	return
}
