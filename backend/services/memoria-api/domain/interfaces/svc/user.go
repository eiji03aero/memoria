package svc

type User interface {
	HasValidStatusForUse(userID string) (ok bool, err error)
}
