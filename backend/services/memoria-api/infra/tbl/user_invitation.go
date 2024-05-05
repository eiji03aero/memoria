package tbl

import (
	"time"

	"memoria-api/domain/model"
	"memoria-api/domain/value"
)

type UserInvitation struct {
	ID          string    `gorm:"column:id"`
	UserID      string    `gorm:"column:user_id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	Type        string    `gorm:"column:type"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t UserInvitation) TableNmae() string {
	return "user_invitations"
}

func (t UserInvitation) ToModel() (ui *model.UserInvitation, err error) {
	uit, err := value.NewUserInvitationType(t.Type)
	if err != nil {
		return
	}

	return model.NewUserInvitation(model.NewUserInvitationDTO{
		ID:          t.ID,
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
		Type:        uit,
	})
}
