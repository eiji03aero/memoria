package registry

import (
	"memoria-api/domain/interfaces"
	"memoria-api/domain/repository"
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
