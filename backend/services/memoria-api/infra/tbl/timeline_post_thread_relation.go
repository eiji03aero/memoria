package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelinePostThreadRelation struct {
	TimelinePostID string    `gorm:"column:timeline_post_id"`
	ThreadID       string    `gorm:"column:thread_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
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
