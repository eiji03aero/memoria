package model

import "time"

type TimelinePostThreadRelation struct {
	TimelinePostID string
	ThreadID       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NewTimelinePostThreadRelationDTO struct {
	TimelinePostID string
	ThreadID       string
}

func NewTimelinePostThreadRelation(dto NewTimelinePostThreadRelationDTO) *TimelinePostThreadRelation {
	return &TimelinePostThreadRelation{
		TimelinePostID: dto.TimelinePostID,
		ThreadID:       dto.ThreadID,
	}
}
