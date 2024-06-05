package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type UserSpaceActivity struct {
	ID          string    `gorm:"column:id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	Type        string    `gorm:"column:type"`
	Data        string    `gorm:"column:data"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t UserSpaceActivity) TableName() string {
	return "user_space_activities"
}

func (t UserSpaceActivity) ToModel() (usa *model.UserSpaceActivity, err error) {
	usa, err = model.NewUserSpaceActivity(model.NewUserSpaceActivityDTO{
		ID:          t.ID,
		UserSpaceID: t.UserSpaceID,
		Type:        t.Type,
		Data:        t.Data,
	})
	if err != nil {
		return
	}

	usa.CreatedAt = t.CreatedAt
	usa.UpdatedAt = t.UpdatedAt
	return
}

func (t *UserSpaceActivity) FromModel(usa *model.UserSpaceActivity) {
	t.ID = usa.ID
	t.UserSpaceID = usa.UserSpaceID
	t.Type = string(usa.Type)
	t.Data = usa.Data
	t.CreatedAt = usa.CreatedAt
	t.UpdatedAt = usa.UpdatedAt
}
