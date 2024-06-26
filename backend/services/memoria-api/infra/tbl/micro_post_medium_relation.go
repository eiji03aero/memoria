package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type MicroPostMediumRelation struct {
	MicroPostID string `gorm:"column:"micro_post_id""`
	MediumID    string `gorm:"column:medium_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t MicroPostMediumRelation) TableName() string {
	return "micro_post_medium_relations"
}

func (t MicroPostMediumRelation) ToModel() (m *model.MicroPostMediumRelation) {
	m = model.NewMicroPostMediumRelation(model.NewMicroPostMediumRelationDTO{
		MicroPostID: t.MicroPostID,
		MediumID:    t.MediumID,
	})

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}
