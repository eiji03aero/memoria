package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelineUnit struct {
	Type      string    `gorm:"column:type"`
	Data      string    `gorm:"column:data"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (t TimelineUnit) ToModel() (m *model.TimelineUnit) {
	m = model.NewTimelineUnit(model.NewTimelineUnitDTO{
		Type: t.Type,
		Data: t.Data,
	})

	m.CreatedAt = t.CreatedAt
	return
}
