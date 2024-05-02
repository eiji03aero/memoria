package value

import (
	"fmt"

	"memoria-api/domain/cerrors"
)

type UserInvitationType string

var (
	UserInvitationType_Signup = UserInvitationType("signup")
	UserInvitationType_Invite = UserInvitationType("invite")
	UserInvitationTypes       = []UserInvitationType{
		UserInvitationType_Signup,
		UserInvitationType_Invite,
	}
)

func IsValidUserInvitationType(t string) bool {
	for _, userInvitationType := range UserInvitationTypes {
		if string(userInvitationType) == t {
			return true
		}
	}
	return false
}

func NewUserInvitationType(ts string) (t UserInvitationType, err error) {
	if !IsValidUserInvitationType(ts) {
		err = cerrors.NewConsistency(fmt.Sprintf("invalid user invitation type: %s", ts))
		return
	}

	t = UserInvitationType(ts)
	return
}
