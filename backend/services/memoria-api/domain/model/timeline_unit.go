package model

import (
	"time"

	"memoria-api/domain/value"
)

type TimelineUnit struct {
	Type      value.TimelineUnitType
	Data      any
	CreatedAt time.Time
}

type NewTimelineUnitDTO struct {
	Type string
	Data any
}

func NewTimelineUnit(dto NewTimelineUnitDTO) (tu *TimelineUnit) {
	t, err := value.NewTimelineUnitType(dto.Type)
	if err != nil {
		return
	}

	return &TimelineUnit{
		Type: t,
		Data: dto.Data,
	}
}
