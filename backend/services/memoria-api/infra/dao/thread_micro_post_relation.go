package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type threadMicroPostRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewThreadMicroPostRelation(db *gorm.DB) repository.ThreadMicroPostRelation {
	return &threadMicroPostRelation[tbl.Thread]{db: db}
}

func (d *threadMicroPostRelation[T]) Create(dto repository.ThreadMicroPostRelationDTO) (tmpr *model.ThreadMicroPostRelation, err error) {
	tmprTbl := &tbl.ThreadMicroPostRelation{
		ThreadID:    dto.ThreadID,
		MicroPostID: dto.MicroPostID,
	}

	err = d.db.Create(tmprTbl).Error
	if err != nil {
		return
	}

	tmpr = tmprTbl.ToModel()
	return
}
