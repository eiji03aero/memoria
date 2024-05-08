package dao

import (
	"errors"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"

	"gorm.io/gorm"
)

type BaseDao[T any] struct{}

type findWithFindOptionDTO struct {
	db         *gorm.DB
	findOption *repository.FindOption
	data       any
	name       string
}

func (d *BaseDao[T]) findWithFindOption(dto findWithFindOptionDTO) (result *gorm.DB, err error) {
	result = dto.db.Scopes(d.scopeByFindOption(dto.findOption)).Find(dto.data)
	err = result.Error
	if isRNF, dmnErr := d.handleResourceNotFound(result.Error, dto.name); isRNF {
		err = dmnErr
		return
	}

	return
}

type findOneWithFindOptionDTO struct {
	db         *gorm.DB
	findOption *repository.FindOption
	data       any
	name       string
}

func (d *BaseDao[T]) findOneWithFindOption(dto findOneWithFindOptionDTO) (result *gorm.DB, err error) {
	result = dto.db.Scopes(d.scopeByFindOption(dto.findOption)).First(dto.data)
	err = result.Error
	if isRNF, dmnErr := d.handleResourceNotFound(result.Error, dto.name); isRNF {
		err = dmnErr
		return
	}

	return
}

func (d *BaseDao[T]) scopeByFindOption(findOption *repository.FindOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := db
		for _, filter := range findOption.Filters {
			query = query.Where(filter.Query, filter.Value)
		}

		if findOption.Order != "" {
			query = query.Order(findOption.Order)
		}

		return query
	}
}

func (d *BaseDao[T]) handleResourceNotFound(err error, name string) (isRNF bool, dmnErr error) {
	isRNF = errors.Is(err, gorm.ErrRecordNotFound)
	if isRNF {
		dmnErr = cerrors.NewResourceNotFound(name)
		return
	}

	return
}
