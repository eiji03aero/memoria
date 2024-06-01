package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type Thread struct {
	ID          string    `gorm:"column:id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t Thread) TableName() string {
	return "threads"
}

func (t Thread) ToModel() (m *model.Thread) {
	m = model.NewThread(model.NewThreadDTO{
		ID:          t.ID,
		UserSpaceID: t.UserSpaceID,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}

func (t *Thread) FromModel(m *model.Thread) {
	t.ID = m.ID
	t.UserSpaceID = m.UserSpaceID
	t.CreatedAt = m.CreatedAt
	t.UpdatedAt = m.UpdatedAt
}
