package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type ThreadMicroPostRelation struct {
	ThreadID    string `gorm:"column:thread_id"`
	MicroPostID string `gorm:"column:micro_post_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t ThreadMicroPostRelation) TableName() string {
	return "thread_micro_post_relations"
}

func (t ThreadMicroPostRelation) ToModel() (m *model.ThreadMicroPostRelation) {
	m = model.NewThreadMicroPostRelation(model.NewThreadMicroPostRelationDTO{
		ThreadID:    t.ThreadID,
		MicroPostID: t.MicroPostID,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}
