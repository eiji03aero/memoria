package dao

import (
	"errors"

	"gorm.io/gorm"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
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
	if err != nil {
		return
	}

	user, err = userTbl.ToModel()
	return
}

func (d *user) Create(dto repository.UserCreateDTO) (err error) {
	userTbl := &tbl.User{
		ID:            dto.ID,
		AccountStatus: dto.AccountStatus,
		Name:          dto.Name,
		Email:         dto.Email,
		PasswordHash:  dto.PasswordHash,
	}

	result := d.db.Create(userTbl)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user")
	}

	return
}

func (d *user) Update(user *model.User) (err error) {
	userTbl := &tbl.User{}
	userTbl.FromModel(user)

	result := d.db.Save(userTbl)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user")
	}

	return
}
