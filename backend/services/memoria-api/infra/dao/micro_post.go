package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type microPost[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewMicroPost(db *gorm.DB) repository.MicroPost {
	return &microPost[tbl.MicroPost]{db: db}
}

func (d *microPost[T]) Find(fOpt *repository.FindOption) (mps []*model.MicroPost, err error) {
	mpTbls := []*tbl.MicroPost{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &mpTbls,
		name:       "micro-post",
	})
	if err != nil {
		return
	}

	mps = make([]*model.MicroPost, 0, len(mpTbls))
	for _, mpTbl := range mpTbls {
		mps = append(mps, mpTbl.ToModel())
	}

	return
}

func (d *microPost[T]) Create(dto repository.MicroPostCreateDTO) (m *model.MicroPost, err error) {
	mTbl := &tbl.MicroPost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
	}

	err = d.db.Create(mTbl).Error
	if err != nil {
		return
	}

	m = mTbl.ToModel()
	return
}
