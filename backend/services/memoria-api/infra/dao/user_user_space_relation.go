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

func (d *userUserSpaceRelation[T]) FindOne(findOption *repository.FindOption) (uusr *model.UserUserSpaceRelation, err error) {
	uusrTbl := &tbl.UserUserSpaceRelation{}
	query := d.ScopeByFindOption(d.db, findOption)
	err = query.First(uusrTbl).Error
	if ok, dmnErr := d.handleResourceNotFound(err, "user user space relation"); ok {
		err = dmnErr
		return
	}
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
