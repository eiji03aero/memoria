package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type timelinePost[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewTimelinePost(db *gorm.DB) repository.TimelinePost {
	return &timelinePost[tbl.TimelinePost]{db: db}
}

func (d *timelinePost[T]) Find(fOpt *repository.FindOption) (tps []*model.TimelinePost, err error) {
	tpTbls := []*tbl.TimelinePost{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &tpTbls,
		name:       "timeline-post",
	})
	if err != nil {
		return
	}

	tps = make([]*model.TimelinePost, 0, len(tpTbls))
	for _, tpTbl := range tpTbls {
		tps = append(tps, tpTbl.ToModel())
	}
	return
}

func (d *timelinePost[T]) Create(dto repository.TimelinePostCreateDTO) (tp *model.TimelinePost, err error) {
	tpTbl := &tbl.TimelinePost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	}

	err = d.db.Create(tpTbl).Error
	if err != nil {
		return
	}

	tp = tpTbl.ToModel()
	return
}
