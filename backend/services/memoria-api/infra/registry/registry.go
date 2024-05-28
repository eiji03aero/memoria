package registry

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/interfaces/svc"
	"memoria-api/domain/service"
	"memoria-api/infra/bgjobivkr"
	"memoria-api/infra/caws"
	"memoria-api/infra/dao"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gorm.io/gorm"
)

type Registry struct {
	DB              *gorm.DB
	awsCfg          aws.Config
	bgjobInvokeChan chan interfaces.BGJobInvokePayload
}

// -------------------- database --------------------
func (r *Registry) BeginTx() {
	r.DB.Begin()
}

func (r *Registry) RollbackTx() {
	r.DB.Rollback()
}

func (r *Registry) CommitTx() {
	r.DB.Commit()
}

func (r *Registry) CloseDB() {
	d, err := r.DB.DB()
	if err != nil {
		panic(err)
	}

	d.Close()
}

// -------------------- bgjob --------------------
func (r *Registry) NewBGJobInvoker() interfaces.BGJobInvoker {
	return bgjobivkr.NewBGJobInvoker(r.bgjobInvokeChan)
}

// -------------------- aws  --------------------
func (r *Registry) NewSESMailer() (interfaces.Mailer, error) {
	return caws.NewSESMailer(r.awsCfg)
}

func (r *Registry) NewS3Client() interfaces.S3Client {
	return caws.NewS3Client(r.awsCfg)
}

// -------------------- repository --------------------
func (r *Registry) NewUserRepository() repository.User {
	return dao.NewUser(r.DB)
}

func (r *Registry) NewUserSpaceRepository() repository.UserSpace {
	return dao.NewUserSpace(r.DB)
}

func (r *Registry) NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation {
	return dao.NewUserUserSpaceRelation(r.DB)
}

func (r *Registry) NewUserInvitationRepository() repository.UserInvitation {
	return dao.NewUserInvitation(r.DB)
}

func (r *Registry) NewAlbumRepository() repository.Album {
	return dao.NewAlbum(r.DB)
}

func (r *Registry) NewUserSpaceAlbumRelationRepository() repository.UserSpaceAlbumRelation {
	return dao.NewUserSpaceAlbumRelation(r.DB)
}

func (r *Registry) NewMediumRepository() repository.Medium {
	return dao.NewMedium(r.DB)
}

func (r *Registry) NewAlbumMediumRelationRepository() repository.AlbumMediumRelation {
	return dao.NewAlbumMediumRelation(r.DB)
}

func (r *Registry) NewUserSpaceActivityRepository() repository.UserSpaceActivity {
	return dao.NewUserSpaceActivity(r.DB)
}

// -------------------- service --------------------
func (r *Registry) NewUserService() svc.User {
	return service.NewUser(service.NewUserDTO{UserRepository: r.NewUserRepository()})
}

func (r *Registry) NewUserInvitationService() svc.UserInvitation {
	return service.NewUserInvitation(service.NewUserInvitationDTO{UserInvitationRepository: r.NewUserInvitationRepository()})
}

func (r *Registry) NewUserSpaceService() svc.UserSpace {
	return service.NewUserSpace(service.NewUserSpaceDTO{UserSpaceRepository: r.NewUserSpaceRepository()})
}

func (r *Registry) NewUserUserSpaceRelationService() svc.UserUserSpaceRelation {
	return service.NewUserUserSpaceRelation(service.NewUserUserSpaceRelationDTO{UserUserSpaceRelationRepository: r.NewUserUserSpaceRelationRepository()})
}

func (r *Registry) NewMediumService() svc.Medium {
	return service.NewMedium(service.NewMediumDTO{S3Client: r.NewS3Client(), MediumRepo: r.NewMediumRepository()})
}

func (r *Registry) NewAlbumService() svc.Album {
	return service.NewAlbum(r)
}

func (r *Registry) NewUserSpaceActivityService() svc.UserSpaceActivity {
	return service.NewUserSpaceActivity(r)
}
