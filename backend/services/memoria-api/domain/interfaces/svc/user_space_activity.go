package svc

type UserSpaceActivity interface {
	CreateInviteUserJoined(dto UserSpaceActivityCreateInviteUserJoined) error
	CreateUserUploadedMedia(dto UserSpaceActivityCreateUserUploadedMedia) error
}

type UserSpaceActivityCreateInviteUserJoined struct {
	UserSpaceID string
	UserID      string
}

type UserSpaceActivityCreateUserUploadedMedia struct {
	UserSpaceID string
	UserID      string
	MediumIDs   []string
}
