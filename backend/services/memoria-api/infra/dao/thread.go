package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type thread[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewThread(db *gorm.DB) repository.Thread {
	return &thread[tbl.Thread]{db: db}
}

func (d *thread[T]) Find(fOpt *repository.FindOption) (ts []*model.Thread, err error) {
	tTbls := []*tbl.Thread{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &tTbls,
		name:       "thread",
	})
	if err != nil {
		return
	}

	ts = make([]*model.Thread, 0, len(tTbls))
	for _, tTbl := range tTbls {
		ts = append(ts, tTbl.ToModel())
	}
	return
}

func (d *thread[T]) FindOne(fOpt *repository.FindOption) (t *model.Thread, err error) {
	tTbl := &tbl.Thread{}

	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       tTbl,
		name:       "thread",
	})

	t = tTbl.ToModel()

	return
}

func (d *thread[T]) Create(dto repository.ThreadCreateDTO) (t *model.Thread, err error) {
	tTbl := &tbl.Thread{
		ID:          dto.ID,
		UserSpaceID: dto.UserSpaceID,
	}

	err = d.db.Create(tTbl).Error
	if err != nil {
		return
	}

	t = tTbl.ToModel()
	return
}
