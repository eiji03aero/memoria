package dao

import (
	"errors"

	"gorm.io/gorm"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/model"
	"memoria-api/domain/repository"
	"memoria-api/infra/tbl"
)

type userSpace struct {
	db *gorm.DB
}

func NewUserSpace(db *gorm.DB) repository.UserSpace {
	return &userSpace{db: db}
}

func (d *userSpace) FindByID(dto repository.UserSpaceFindByIDDTO) (userSpace *model.UserSpace, err error) {
	userSpaceTbl := &tbl.UserSpace{ID: dto.ID}
	err = d.db.First(userSpaceTbl).Error
	userSpace = userSpaceTbl.ToModel()
	return
}

func (d *userSpace) Create(dto repository.UserSpaceCreateDTO) (err error) {
	userSpace := &tbl.UserSpace{
		ID:   dto.ID,
		Name: dto.Name,
	}

	result := d.db.Create(userSpace)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user space")
	}

	return nil
}
