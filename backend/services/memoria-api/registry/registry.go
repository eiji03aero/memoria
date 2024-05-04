package registry

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/service"
	"memoria-api/infra/caws"
	"memoria-api/infra/dao"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gorm.io/gorm"
)

type Registry interface {
	// database
	BeginTx()
	RollbackTx()
	CommitTx()
	// mailer
	NewSESMailer() (interfaces.Mailer, error)
	// repository
	NewUserRepository() repository.User
	NewUserSpaceRepository() repository.UserSpace
	NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation
	NewUserInvitationRepository() repository.UserInvitation
	// service
	NewUserService() *service.User
	NewUserInvitationService() *service.UserInvitation
	NewUserSpaceService() *service.UserSpace
	NewUserUserSpaceRelationService() *service.UserUserSpaceRelation
}

type registry struct {
	db     *gorm.DB
	awsCfg aws.Config
}

type NewRegistryDTO struct {
	DB     *gorm.DB
	AWSCfg aws.Config
}

func NewRegistry(dto NewRegistryDTO) Registry {
	return &registry{
		db:     dto.DB,
		awsCfg: dto.AWSCfg,
	}
}

// -------------------- database --------------------
func (r *registry) BeginTx() {
	r.db.Begin()
}

func (r *registry) RollbackTx() {
	r.db.Rollback()
}

func (r *registry) CommitTx() {
	r.db.Commit()
}

// -------------------- mailer --------------------
func (r *registry) NewSESMailer() (interfaces.Mailer, error) {
	return caws.NewSESMailer(r.awsCfg)
}

// -------------------- repository --------------------
func (r *registry) NewUserRepository() repository.User {
	return dao.NewUser(r.db)
}

func (r *registry) NewUserSpaceRepository() repository.UserSpace {
	return dao.NewUserSpace(r.db)
}

func (r *registry) NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation {
	return dao.NewUserUserSpaceRelation(r.db)
}

func (r *registry) NewUserInvitationRepository() repository.UserInvitation {
	return dao.NewUserInvitation(r.db)
}

// -------------------- service --------------------
func (r *registry) NewUserService() *service.User {
	return service.NewUser(service.NewUserDTO{UserRepository: r.NewUserRepository()})
}

func (r *registry) NewUserInvitationService() *service.UserInvitation {
	return service.NewUserInvitation(service.NewUserInvitationDTO{UserInvitationRepository: r.NewUserInvitationRepository()})
}

func (r *registry) NewUserSpaceService() *service.UserSpace {
	return service.NewUserSpace(service.NewUserSpaceDTO{UserSpaceRepository: r.NewUserSpaceRepository()})
}

func (r *registry) NewUserUserSpaceRelationService() *service.UserUserSpaceRelation {
	return service.NewUserUserSpaceRelation(service.NewUserUserSpaceRelationDTO{UserUserSpaceRelationRepository: r.NewUserUserSpaceRelationRepository()})
}
