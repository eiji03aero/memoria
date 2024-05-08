package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type userUserSpaceRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUserUserSpaceRelation(db *gorm.DB) repository.UserUserSpaceRelation {
	return &userUserSpaceRelation[tbl.UserUserSpaceRelation]{db: db}
}

func (d *userUserSpaceRelation[T]) Find(findOption *repository.FindOption) (uusrs []*model.UserUserSpaceRelation, err error) {
	uusrTbls := []*tbl.UserUserSpaceRelation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &uusrTbls,
		name:       "user-user-space-relation",
	})
	if err != nil {
		return
	}

	uusrs = make([]*model.UserUserSpaceRelation, 0, len(uusrTbls))
	for _, uusrTbl := range uusrTbls {
		uusrs = append(uusrs, uusrTbl.ToModel())
	}
	return
}

func (d *userUserSpaceRelation[T]) FindOne(findOption *repository.FindOption) (uusr *model.UserUserSpaceRelation, err error) {
	uusrTbl := tbl.UserUserSpaceRelation{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &uusrTbl,
		name:       "user-user-space-relation",
	})
	if err != nil {
		return
	}

	uusr = uusrTbl.ToModel()
	return
}

func (d *userUserSpaceRelation[T]) Create(dto repository.UserUserSpaceRelationCreateDTO) (err error) {
	uusr := &tbl.UserUserSpaceRelation{
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	}

	result := d.db.Create(uusr)
	err = result.Error
	return
}
