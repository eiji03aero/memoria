package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type userSpaceActivity[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewUserSpaceActivity(db *gorm.DB) repository.UserSpaceActivity {
	return &userSpaceActivity[tbl.UserSpaceActivity]{db: db}
}

func (d *userSpaceActivity[T]) Find(findOption *repository.FindOption) (usas []*model.UserSpaceActivity, err error) {
	usaTbls := []*tbl.UserSpaceActivity{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &usaTbls,
		name:       "user-space-activity",
	})
	if err != nil {
		return
	}

	usas = make([]*model.UserSpaceActivity, 0, len(usaTbls))
	for _, usaTbl := range usaTbls {
		usa, e := usaTbl.ToModel()
		if e != nil {
			err = e
			return
		}

		usas = append(usas, usa)
	}
	return
}

func (d *userSpaceActivity[T]) FindOne(findOption *repository.FindOption) (usa *model.UserSpaceActivity, err error) {
	usaTbl := &tbl.UserSpaceActivity{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: findOption,
		data:       &usaTbl,
		name:       "user-space-activity",
	})
	if err != nil {
		return
	}

	usa, err = usaTbl.ToModel()
	if err != nil {
		return
	}

	return
}

func (d *userSpaceActivity[T]) FindOneByID(id string) (usa *model.UserSpaceActivity, err error) {
	return d.FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	})
}

func (d *userSpaceActivity[T]) Create(dto repository.UserSpaceActivityCreateDTO) (usa *model.UserSpaceActivity, err error) {
	usaTbl := &tbl.UserSpaceActivity{
		ID:          dto.ID,
		UserSpaceID: dto.UserSpaceID,
		Type:        string(dto.Type),
		Data:        dto.Data,
	}

	err = d.db.Create(usaTbl).Error
	if err != nil {
		return
	}

	usa, err = usaTbl.ToModel()
	return
}
