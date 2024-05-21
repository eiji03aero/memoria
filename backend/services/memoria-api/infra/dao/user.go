package dao

import (
	"errors"

	"gorm.io/gorm"

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

func (d *user[T]) Find(findOption *repository.FindOption) (users []*model.User, err error) {
	userTbls := []*tbl.User{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &userTbls,
		name:       "user",
	})
	if err != nil {
		return
	}

	users = make([]*model.User, 0, len(userTbls))
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

func (d *user[T]) FindOne(findOption *repository.FindOption) (user *model.User, err error) {
	userTbl := tbl.User{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &userTbl,
		name:       "user",
	})
	if err != nil {
		return
	}

	user, err = userTbl.ToModel()
	return
}

func (d *user[T]) FindByID(userID string) (user *model.User, err error) {
	return d.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: userID},
		},
	})
}

func (d *user[T]) Exists(findOpt *repository.FindOption) (exists bool, err error) {
	return d.exists(existsDTO{
		db:         d.db,
		findOption: findOpt,
		data:       &tbl.User{},
		name:       "user",
	})
}

func (d *user[T]) EmailExistsInUserSpace(userSpaceID string, email string) (exists bool, err error) {
	userTbl := &tbl.User{}
	err = d.db.Table("users").
		Select("users.email").
		Joins("join user_user_space_relations uusr on uusr.user_id = users.id").
		Where("uusr.user_space_id = ?", userSpaceID).
		Where("users.email = ?", email).
		First(userTbl).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		exists = false
		err = nil
		return
	}
	if err != nil {
		exists = false
		return
	}

	exists = true
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

	return
}

func (d *user[T]) Update(user *model.User) (err error) {
	userTbl := &tbl.User{}
	userTbl.FromModel(user)

	result := d.db.Save(userTbl)
	err = result.Error
	if isRNF, dmnErr := d.handleResourceNotFound(err, "user"); isRNF {
		err = dmnErr
		return
	}

	return
}
