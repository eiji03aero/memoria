package dao

import (
	"errors"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"

	"gorm.io/gorm"
)

type BaseDao[T any] struct{}

func (d *BaseDao[T]) ScopeByFindOption(db *gorm.DB, findOption *repository.FindOption) *gorm.DB {
	query := db
	for _, filter := range findOption.Filters {
		query = db.Where(filter.Query, filter.Value)
	}
	return query
}

func (d *BaseDao[T]) handleResourceNotFound(err error, name string) (isRNF bool, dmnErr error) {
	isRNF = errors.Is(err, gorm.ErrRecordNotFound)
	if isRNF {
		dmnErr = cerrors.NewResourceNotFound(name)
		return
	}

	return
}
