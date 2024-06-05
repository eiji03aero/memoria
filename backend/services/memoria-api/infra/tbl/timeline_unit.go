package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelineUnit struct {
	ID        string    `gorm:"column:id"`
	Type      string    `gorm:"column:type"`
	Data      string    `gorm:"column:data"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (t TimelineUnit) ToModel() (m *model.TimelineUnit, err error) {
	m, err = model.NewTimelineUnit(model.NewTimelineUnitDTO{
		ID:   t.ID,
		Type: t.Type,
		Data: t.Data,
	})
	if err != nil {
		return
	}

	m.CreatedAt = t.CreatedAt
	return
}
