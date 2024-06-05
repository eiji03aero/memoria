package model

import (
	"time"

	"memoria-api/domain/value"
)

type TimelineUnit struct {
	ID        string
	Type      value.TimelineUnitType
	Data      any
	CreatedAt time.Time
}

type NewTimelineUnitDTO struct {
	ID   string
	Type string
	Data any
}

func NewTimelineUnit(dto NewTimelineUnitDTO) (tu *TimelineUnit, err error) {
	t, err := value.NewTimelineUnitType(dto.Type)
	if err != nil {
		return
	}

	tu = &TimelineUnit{
		ID:   dto.ID,
		Type: t,
		Data: dto.Data,
	}
	return
}

func (tu *TimelineUnit) ParseData_TimelineUnit() {
}

type TimelineUnits []*TimelineUnit

func (tus TimelineUnits) Len() int {
	return len(tus)
}

func (tus TimelineUnits) Less(i, j int) bool {
	return tus[i].ID < tus[j].ID
}

func (tus TimelineUnits) Swap(i, j int) {
	tus[i], tus[j] = tus[j], tus[i]
}
