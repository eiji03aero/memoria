package dao

import (
	"errors"

	"gorm.io/gorm"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type user[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &user[tbl.User]{db: db}
}

func (d *user[T]) Find(dto repository.UserFindDTO) (users []*model.User, err error) {
	userTbls := []tbl.User{}

	query := d.ScopeByFindOption(ScopeByFindOptionDTO{db: d.db, findOption: dto.FindOption})
	result := query.Find(&userTbls)
	err = result.Error
	if err != nil {
		return
	}

	users = make([]*model.User, len(userTbls))
	for _, userTbl := range userTbls {
		user, e := userTbl.ToModel()
		if e != nil {
			err = e
			return
		}

		users = append(users, user)
	}

	return
}

func (d *user[T]) FindByID(dto repository.UserFindByIDDTO) (user *model.User, err error) {
	userTbl := &tbl.User{ID: dto.ID}
	err = d.db.First(userTbl).Error
	if err != nil {
		return
	}

	user, err = userTbl.ToModel()
	return
}

func (d *user[T]) Create(dto repository.UserCreateDTO) (err error) {
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

func (d *user[T]) Update(user *model.User) (err error) {
	userTbl := &tbl.User{}
	userTbl.FromModel(user)

	result := d.db.Save(userTbl)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user")
	}

	return
}
