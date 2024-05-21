package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type medium[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewMedium(db *gorm.DB) repository.Medium {
	return &medium[tbl.Medium]{db: db}
}

func (d *medium[T]) Find(findOption *repository.FindOption) (media []*model.Medium, err error) {
	mediumTbls := []*tbl.Medium{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &mediumTbls,
		name:       "medium",
	})
	if err != nil {
		return
	}

	media = make([]*model.Medium, 0, len(mediumTbls))
	for _, mediumTbl := range mediumTbls {
		media = append(media, mediumTbl.ToModel())
	}
	return
}

func (d *medium[T]) FindOne(findOption *repository.FindOption) (medium *model.Medium, err error) {
	mediumTbl := &tbl.Medium{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &mediumTbl,
		name:       "medium",
	})
	if err != nil {
		return
	}

	medium = mediumTbl.ToModel()
	return
}

func (d *medium[T]) Create(dto repository.MediumCreateDTO) (medium *model.Medium, err error) {
	mediumTbl := &tbl.Medium{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Name:        dto.Name,
		Extension:   dto.Extension,
	}

	err = d.db.Create(mediumTbl).Error
	if err != nil {
		return
	}

	medium = mediumTbl.ToModel()
	return
}

func (d *medium[T]) Delete(dto repository.MediumDeleteDTO) (err error) {
	mediumTbl := &tbl.Medium{
		ID: dto.ID,
	}

	err = d.db.Delete(mediumTbl).Error
	return
}
