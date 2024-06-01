package dao

import (
	"gorm.io/gorm"

	"memoria-api/domain/interfaces/repository"
	"memoria-api/domain/model"
	"memoria-api/infra/tbl"
)

type timelinePostThreadRelation[T any] struct {
	BaseDao[T]
	db *gorm.DB
}

func NewTimelinePostThreadRelation(db *gorm.DB) repository.TimelinePostThreadRelation {
	return &timelinePostThreadRelation[tbl.TimelinePostThreadRelation]{db: db}
}

func (d *timelinePostThreadRelation[T]) Find(fOpt *repository.FindOption) (tptrs []*model.TimelinePostThreadRelation, err error) {
	tptrTbls := []*tbl.TimelinePostThreadRelation{}
	_, err = d.findWithFindOption(findWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &tptrTbls,
		name:       "timeline-post-thread-relation",
	})
	if err != nil {
		return
	}

	tptrs = make([]*model.TimelinePostThreadRelation, 0, len(tptrTbls))
	for _, tptrTbl := range tptrTbls {
		tptrs = append(tptrs, tptrTbl.ToModel())
	}
	return
}

func (d *timelinePostThreadRelation[T]) Create(dto repository.TimelinePostThreadRelationCreateDTO) (tp *model.TimelinePostThreadRelation, err error) {
	tptrTbl := &tbl.TimelinePostThreadRelation{
		TimelinePostID: dto.TimelinePostID,
		ThreadID:       dto.ThreadID,
	}

	err = d.db.Create(tptrTbl).Error
	if err != nil {
		return
	}

	tp = tptrTbl.ToModel()
	return
}
