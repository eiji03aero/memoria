package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelinePostThreadRelation struct {
	TimelinePostID string
	ThreadID       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (t TimelinePostThreadRelation) TableName() string {
	return "timeline_post_thread_relations"
}

func (t *TimelinePostThreadRelation) ToModel() (m *model.TimelinePostThreadRelation) {
	m = model.NewTimelinePostThreadRelation(model.NewTimelinePostThreadRelationDTO{
		TimelinePostID: t.TimelinePostID,
		ThreadID:       t.ThreadID,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}
