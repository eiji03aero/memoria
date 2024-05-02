package value

import (
	"fmt"

	"memoria-api/domain/cerrors"
)

type UserAccountStatus string

var (
	UserAccountStatus_Invited   = UserAccountStatus("invited")
	UserAccountStatus_Confirmed = UserAccountStatus("confirmed")
	UserAccountStatuses         = []UserAccountStatus{
		UserAccountStatus_Invited,
		UserAccountStatus_Confirmed,
	}
)

func IsValidUserAccountStatus(status string) bool {
	for _, userAccountStatus := range UserAccountStatuses {
		if string(userAccountStatus) == status {
			return true
		}
	}
	return false
}

func NewUserAccountStatus(status string) (sts UserAccountStatus, err error) {
	if !IsValidUserAccountStatus(status) {
		err = cerrors.NewConsistency(fmt.Sprintf("invalid user account status: %s", status))
		return
	}

	sts = UserAccountStatus(status)
	return
}
