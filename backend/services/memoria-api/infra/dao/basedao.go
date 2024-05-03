package dao

import (
	"memoria-api/domain/interfaces/repository"

	"gorm.io/gorm"
)

type BaseDao[T any] struct{}

type ScopeByFindOptionDTO struct {
	db         *gorm.DB
	findOption *repository.FindOption
}

func (d *BaseDao[T]) ScopeByFindOption(dto ScopeByFindOptionDTO) *gorm.DB {
	db := dto.db.Debug()
	for _, filter := range dto.findOption.Filters {
		db = db.Where(filter.Query, filter.Value)
	}
	return db
}
