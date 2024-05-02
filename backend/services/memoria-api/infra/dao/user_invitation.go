package dao

import (
	"errors"

	"memoria-api/domain/cerrors"
	"memoria-api/domain/model"
	"memoria-api/domain/repository"
	"memoria-api/infra/tbl"

	"gorm.io/gorm"
)

type userInvitation struct {
	db *gorm.DB
}

func NewUserInvitation(db *gorm.DB) repository.UserInvitation {
	return &userInvitation{db: db}
}

func (d *userInvitation) FindByID(dto repository.UserInvitationFindByIDDTO) (ui *model.UserInvitation, err error) {
	uiTbl := &tbl.UserInvitation{ID: dto.ID}
	err = d.db.First(uiTbl).Error
	if err != nil {
		return
	}

	ui, err = uiTbl.ToModel()
	return
}

func (d *userInvitation) Create(dto repository.UserInvitationCreateDTO) (err error) {
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
