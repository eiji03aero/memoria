package value

import (
	"fmt"

	"memoria-api/domain/cerrors"
)

type UserSpaceActivityType string

var (
	UserSpaceActivityType_InviteUserJoined  = UserSpaceActivityType("invite-user-joined")
	UserSpaceActivityType_UserUploadedMedia = UserSpaceActivityType("user-uploaded-media")
	UserSpaceActivityTypes                  = []UserSpaceActivityType{
		UserSpaceActivityType_InviteUserJoined,
		UserSpaceActivityType_UserUploadedMedia,
	}
)

func IsValidUserSpaceActivityType(t string) bool {
	for _, usat := range UserSpaceActivityTypes {
		if string(usat) == t {
			return true
		}
	}
	return false
}

func NewUserSpaceActivityType(ts string) (t UserSpaceActivityType, err error) {
	if !IsValidUserSpaceActivityType(ts) {
		err = cerrors.NewConsistency(fmt.Sprintf("invalid user space activity type: %s", ts))
		return
	}

	t = UserSpaceActivityType(ts)
	return
}
