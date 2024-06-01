package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type TimelinePost struct {
	ID          string    `gorm:"column:id"`
	UserID      string    `gorm:"column:user_id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t TimelinePost) TableName() string {
	return "timeline_posts"
}

func (t TimelinePost) ToModel() (m *model.TimelinePost) {
	m = model.NewTimelinePost(model.NewTimelinePostDTO{
		ID:          t.ID,
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}

func (t *TimelinePost) FromModel(m *model.TimelinePost) {
	t.ID = m.ID
	t.UserID = m.UserID
	t.UserSpaceID = m.UserSpaceID
	t.CreatedAt = m.CreatedAt
	t.UpdatedAt = m.UpdatedAt
}
