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

func (d *userInvitation[T]) Find(findOption *repository.FindOption) (uis []*model.UserInvitation, err error) {
	uiTbls := []tbl.UserInvitation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &uiTbls,
		name:       "user-invitation",
	})
	if err != nil {
		return
	}

	uis = make([]*model.UserInvitation, 0, len(uiTbls))
	for _, uiTbl := range uiTbls {
		ui, e := uiTbl.ToModel()
		if err != nil {
			err = e
			return
		}
		uis = append(uis, ui)
	}
	return
}

func (d *userInvitation[T]) FindOne(findOption *repository.FindOption) (ui *model.UserInvitation, err error) {
	uiTbl := tbl.UserInvitation{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &uiTbl,
		name:       "user-invitation",
	})
	if err != nil {
		return
	}

	ui, err = uiTbl.ToModel()
	return
}

func (d *userInvitation[T]) Create(dto repository.UserInvitationCreateDTO) (err error) {
	ui := &tbl.UserInvitation{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Type:        dto.Type,
	}

	result := d.db.Create(ui)
	err = result.Error
	return
}
