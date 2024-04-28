package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/repository"
	"memoria-api/infra/tbl"
)

type userUserSpaceRelation struct {
	db *gorm.DB
}

func NewUserUserSpaceRelation(db *gorm.DB) repository.UserUserSpaceRelation {
	return &userUserSpaceRelation{db: db}
}

func (d *userUserSpaceRelation) Create(dto repository.UserUserSpaceRelationCreateDTO) (err error) {
	uusr := &tbl.UserUserSpaceRelation{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	}

	result := d.db.Create(uusr)
	err = result.Error

	return
}
