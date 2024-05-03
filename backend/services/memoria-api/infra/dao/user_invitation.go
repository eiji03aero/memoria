package dao

import (
	"errors"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"

	"gorm.io/gorm"
)

type userInvitation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUserInvitation(db *gorm.DB) repository.UserInvitation {
	return &userInvitation[tbl.UserInvitation]{db: db}
}

func (d *userInvitation[T]) FindByID(dto repository.UserInvitationFindByIDDTO) (ui *model.UserInvitation, err error) {
	uiTbl := &tbl.UserInvitation{ID: dto.ID}
	err = d.db.First(uiTbl).Error
	if err != nil {
		return
	}

	ui, err = uiTbl.ToModel()
	return
}

func (d *userInvitation[T]) Create(dto repository.UserInvitationCreateDTO) (err error) {
	ui := &tbl.UserInvitation{
		ID:     dto.ID,
		UserID: dto.UserID,
		Type:   dto.Type,
	}

	result := d.db.Create(ui)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = cerrors.NewResourceNotFound("user invitation")
	}

	return
}