package model

import (
	"encoding/json"
	"time"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/value"
)

type UserSpaceActivity struct {
	ID          string
	UserSpaceID string
	Type        value.UserSpaceActivityType
	Data        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewUserSpaceActivityDTO struct {
	ID          string
	UserSpaceID string
	Type        string
	Data        string
}

func NewUserSpaceActivity(dto NewUserSpaceActivityDTO) (usa *UserSpaceActivity, err error) {
	t, err := value.NewUserSpaceActivityType(dto.Type)
	if err != nil {
		return
	}

	usa = &UserSpaceActivity{
		ID:          dto.ID,
		UserSpaceID: dto.UserSpaceID,
		Type:        t,
		Data:        dto.Data,
	}
	return
}

// -------------------- InviteUserJoined --------------------
type UserSpaceActivityData_InviteUserJoined struct {
	UserID string `json:"user_id"`
}

func (u *UserSpaceActivity) ParseData_InviteUserJoined() (d *UserSpaceActivityData_InviteUserJoined, err error) {
	if u.Type != value.UserSpaceActivityType_InviteUserJoined {
		err = cerrors.NewConsistency("ParseData_InviteUserJoined called with wrong type: " + string(u.Type))
		return
	}

	d = &UserSpaceActivityData_InviteUserJoined{}
	err = json.Unmarshal([]byte(u.Data), d)
	return
}

// -------------------- UserUploadedMedia --------------------
type UserSpaceActivityData_UserUploadedMedia struct {
	UserID    string   `json:"user_id"`
	MediumIDs []string `json:"medium_ids"`
}

func (u *UserSpaceActivity) ParseData_UserUploadedMedia() (d *UserSpaceActivityData_UserUploadedMedia, err error) {
	if u.Type != value.UserSpaceActivityType_UserUploadedMedia {
		err = cerrors.NewConsistency("ParseData_UserUploadedMedia called with wrong type: " + string(u.Type))
		return
	}

	d = &UserSpaceActivityData_UserUploadedMedia{}
	err = json.Unmarshal([]byte(u.Data), d)
	return
}
