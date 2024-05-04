package dao

import (
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

func (d *userInvitation[T]) FindOne(findOption *repository.FindOption) (ui *model.UserInvitation, err error) {
	uiTbl := &tbl.UserInvitation{}
	query := d.ScopeByFindOption(d.db, findOption)
	err = query.First(uiTbl).Error
	if ok, dmnErr := d.handleResourceNotFound(err, "user invitation"); ok {
		err = dmnErr
		return
	}
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
	return
}
