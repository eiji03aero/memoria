package interfaces

type Context interface {
	GetUserID() string
	GetUserSpaceID() string
}
