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

type existsDTO struct {
	db         *gorm.DB
	findOption *repository.FindOption
	data       any
	name       string
}

func (d *BaseDao[T]) exists(dto existsDTO) (exists bool, err error) {
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         dto.db,
		findOption: dto.findOption,
		data:       dto.data,
		name:       dto.name,
	})
	if errors.As(err, &cerrors.ResourceNotFound{}) {
		exists = false
		err = nil
		return
	}

	exists = true
	return
}

func (d *BaseDao[T]) scopeByFindOption(findOption *repository.FindOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := db

		for _, filter := range findOption.Filters {
			query = query.Where(filter.Query, filter.Value)
		}

		for _, join := range findOption.Joins {
			query = query.Joins(join.Query)
		}

		if findOption.Order != "" {
			query = query.Order(findOption.Order)
		}

		if findOption.Offset != nil && findOption.Limit != nil {
			query = query.Offset(*findOption.Offset).Limit(*findOption.Limit)
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
