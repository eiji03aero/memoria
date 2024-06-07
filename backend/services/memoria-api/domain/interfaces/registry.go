package interfaces

import (
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
)

type Registry interface {
	// database
	BeginTx()
	RollbackTx()
	CommitTx()
	CloseDB()
	// tools
	NewLogger() Logger
	// bgjob
	NewBGJobInvoker() BGJobInvoker
	// aws
	NewSESMailer() (Mailer, error)
	NewS3Client() S3Client
	// repository
	NewUserRepository() repository.User
	NewUserSpaceRepository() repository.UserSpace
	NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation
	NewUserInvitationRepository() repository.UserInvitation
	NewAlbumRepository() repository.Album
	NewUserSpaceAlbumRelationRepository() repository.UserSpaceAlbumRelation
	NewMediumRepository() repository.Medium
	NewAlbumMediumRelationRepository() repository.AlbumMediumRelation
	NewUserSpaceActivityRepository() repository.UserSpaceActivity
	NewTimelineRepository() repository.Timeline
	NewMicroPostMediumRelationRepository() repository.MicroPostMediumRelation
	NewMicroPostRepository() repository.MicroPost
	NewThreadMicroPostRelationRepository() repository.ThreadMicroPostRelation
	NewThreadRepository() repository.Thread
	NewTimelinePostThreadRelationRepository() repository.TimelinePostThreadRelation
	NewTimelinePostRepository() repository.TimelinePost
	// service
	NewUserService() svc.User
	NewUserInvitationService() svc.UserInvitation
	NewUserSpaceService() svc.UserSpace
	NewUserUserSpaceRelationService() svc.UserUserSpaceRelation
	NewMediumService() svc.Medium
	NewAlbumService() svc.Album
	NewUserSpaceActivityService() svc.UserSpaceActivity
	NewMicroPostService() svc.MicroPost
	NewThreadService() svc.Thread
	NewTimelineService() svc.Timeline
}
