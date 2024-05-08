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

func (d *userSpace[T]) Find(findOption *repository.FindOption) (userSpaces []*model.UserSpace, err error) {
	userSpaceTbls := []*tbl.UserSpace{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &userSpaceTbls,
		name:       "user-space",
	})
	if err != nil {
		return
	}

	userSpaces = make([]*model.UserSpace, 0, len(userSpaceTbls))
	for _, userSpaceTbl := range userSpaceTbls {
		userSpaces = append(userSpaces, userSpaceTbl.ToModel())
	}
	return
}

func (d *userSpace[T]) FindOne(findOption *repository.FindOption) (userSpace *model.UserSpace, err error) {
	userSpaceTbl := tbl.UserSpace{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &userSpaceTbl,
		name:       "user-space",
	})
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
