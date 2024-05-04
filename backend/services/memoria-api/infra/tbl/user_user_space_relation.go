package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type UserUserSpaceRelation struct {
	UserID      string    `gorm:"column:user_id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t UserUserSpaceRelation) TableName() string {
	return "user_user_space_relations"
}

func (t UserUserSpaceRelation) ToModel() *model.UserUserSpaceRelation {
	return model.NewUserUserSpaceRelation(model.NewUserUserSpaceRelationDTO{
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	})
}
