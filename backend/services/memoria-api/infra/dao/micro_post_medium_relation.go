package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type microPostMediumRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewMicroPostMediumRelation(db *gorm.DB) repository.MicroPostMediumRelation {
	return &microPostMediumRelation[tbl.MicroPostMediumRelation]{db: db}
}

func (d *microPostMediumRelation[T]) Find(fOpt *repository.FindOption) (mpmrs []*model.MicroPostMediumRelation, err error) {
	mpmrTbls := []*tbl.MicroPostMediumRelation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &mpmrTbls,
		name:       "micro-post",
	})
	if err != nil {
		return
	}

	mpmrs = make([]*model.MicroPostMediumRelation, 0, len(mpmrTbls))
	for _, mpmrTbl := range mpmrTbls {
		mpmrs = append(mpmrs, mpmrTbl.ToModel())
	}
	return
}

func (d *microPostMediumRelation[T]) Create(dto repository.MicroPostMediumRelationCreateDTO) (mpmr *model.MicroPostMediumRelation, err error) {
	mpmrTbl := &tbl.MicroPostMediumRelation{
		MicroPostID: dto.MicroPostID,
		MediumID:    dto.MediumID,
	}

	err = d.db.Create(mpmrTbl).Error
	if err != nil {
		return
	}

	mpmr = mpmrTbl.ToModel()
	return
}
