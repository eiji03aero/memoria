package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type userSpace[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUserSpace(db *gorm.DB) repository.UserSpace {
	return &userSpace[tbl.UserSpace]{db: db}
}

func (d *userSpace[T]) FindOne(findOption *repository.FindOption) (userSpace *model.UserSpace, err error) {
	userSpaceTbl := &tbl.UserSpace{}
	query := d.ScopeByFindOption(d.db, findOption)
	err = query.First(userSpaceTbl).Error
	if ok, dmnErr := d.handleResourceNotFound(err, "user space"); ok {
		err = dmnErr
		return
	}
	if err != nil {
		return
	}

	userSpace = userSpaceTbl.ToModel()
	return
}

func (d *userSpace[T]) Create(dto repository.UserSpaceCreateDTO) (err error) {
	userSpace := &tbl.UserSpace{
		ID:   dto.ID,
		Name: dto.Name,
	}

	result := d.db.Create(userSpace)
	err = result.Error
	return nil
}
