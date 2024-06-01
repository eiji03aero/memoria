package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type timeline[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewTimeline(db *gorm.DB) repository.Timeline {
	return &timeline[tbl.TimelineUnit]{db: db}
}

type TimelineFindMergedUnit struct {
	Type string `gorm:"column:type"`
}

func (d *timeline[T]) Find(fOpt *repository.FindOption) (tus []*model.TimelineUnit, err error) {
	// get merged result
	//   - get the orders
	//   - get the position
	//   - get the types
	// iterate through the previous result
	//   - for timeline post
	//   - for user space activity

	return
}
