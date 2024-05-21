package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type Medium struct {
	ID          string    `gorm:"column:id"`
	UserID      string    `gorm:"column:user_id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	Name        string    `gorm:"column:name"`
	Extension   string    `gorm:"column:extension"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t Medium) TableName() string {
	return "media"
}

func (t Medium) ToModel() (medium *model.Medium) {
	medium = model.NewMedium(model.NewMediumDTO{
		ID:          t.ID,
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
		Name:        t.Name,
		Extension:   t.Extension,
	})

	medium.CreatedAt = t.CreatedAt
	medium.UpdatedAt = t.UpdatedAt
	return
}

func (t *Medium) FromModel(medium *model.Medium) {
	t.ID = medium.ID
	t.Name = medium.Name
	t.UserID = medium.UserID
	t.UserSpaceID = medium.UserSpaceID
	t.CreatedAt = medium.CreatedAt
	t.UpdatedAt = medium.UpdatedAt
}
