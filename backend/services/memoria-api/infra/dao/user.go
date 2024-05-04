package dao

import (
	"log"

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
	userTbls := []tbl.User{}

	query := d.ScopeByFindOption(d.db, findOption)
	result := query.Find(&userTbls)
	err = result.Error
	if ok, dmnErr := d.handleResourceNotFound(err, "user"); ok {
		err = dmnErr
		return
	}
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

func (d *user[T]) FindOne(findOption *repository.FindOption) (user *model.User, err error) {
	userTbl := &tbl.User{}
	query := d.ScopeByFindOption(d.db, findOption)
	err = query.First(userTbl).Error
	log.Println("going in dao user", err)
	if ok, dmnErr := d.handleResourceNotFound(err, "user"); ok {
		err = dmnErr
		return
	}
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

	return
}

func (d *user[T]) Update(user *model.User) (err error) {
	userTbl := &tbl.User{}
	userTbl.FromModel(user)

	result := d.db.Save(userTbl)
	err = result.Error
	if ok, dmnErr := d.handleResourceNotFound(err, "user"); ok {
		err = dmnErr
		return
	}

	return
}
