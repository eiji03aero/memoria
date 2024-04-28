package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type UserSpace struct {
	ID        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (t UserSpace) TableName() string {
	return "user_spaces"
}

func (t UserSpace) ToModel() *model.UserSpace {
	return model.NewUserSpace(model.NewUserSpaceDTO{
		ID:   t.ID,
		Name: t.Name,
	})
}
