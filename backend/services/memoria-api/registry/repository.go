package registry

import (
	"memoria-api/domain/repository"
	"memoria-api/infra/dao"

	"gorm.io/gorm"
)

type Registry interface {
	BeginTx()
	RollbackTx()
	CommitTx()
	NewUserRepository() repository.User
	NewUserSpaceRepository() repository.UserSpace
	NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation
}

type registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{
		db: db,
	}
}

func (r *registry) BeginTx() {
	r.db.Begin()
}

func (r *registry) RollbackTx() {
	r.db.Rollback()
}

func (r *registry) CommitTx() {
	r.db.Commit()
}

func (r *registry) NewUserRepository() repository.User {
	return dao.NewUser(r.db)
}

func (r *registry) NewUserSpaceRepository() repository.UserSpace {
	return dao.NewUserSpace(r.db)
}

func (r *registry) NewUserUserSpaceRelationRepository() repository.UserUserSpaceRelation {
	return dao.NewUserUserSpaceRelation(r.db)
}
