package dao

import (
	"errors"

	"gorm.io/gorm"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/model"
	"memoria-api/domain/repository"
	"memoria-api/infra/tbl"
)

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &user{db: db}
}

func (d *user) FindByID(dto repository.UserFindByIDDTO) (user *model.User, err error) {
	userTbl := &tbl.User{ID: dto.ID}
	err = d.db.First(userTbl).Error
	user = userTbl.ToModel()
	return
}

func (d *user) Create(dto repository.UserCreateDTO) (err error) {
	user := &tbl.User{
		ID:           dto.ID,
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: dto.PasswordHash,
	}

	result := d.db.Create(user)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user")
	}

	return
}
