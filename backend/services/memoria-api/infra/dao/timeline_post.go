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
		tp, e := tpTbl.ToModel()
		if e != nil {
			err = e
			return
		}

		tps = append(tps, tp)
	}
	return
}

func (d *timelinePost[T]) FindOne(fOpt *repository.FindOption) (tp *model.TimelinePost, err error) {
	tpTbl := &tbl.TimelinePost{}
	_, err = d.findOneWithFindOption(findOneWithFindOptionDTO{
		db:         d.db,
		findOption: fOpt,
		data:       &tpTbl,
		name:       "timeline-post",
	})
	if err != nil {
		return
	}

	tp, err = tpTbl.ToModel()
	if err != nil {
		return
	}

	return
}

func (d *timelinePost[T]) FindOneByID(id string, fOpt *repository.FindOption) (tp *model.TimelinePost, err error) {
	findOption := &repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "id = ?", Value: id},
		},
	}
	findOption.Merge(fOpt)
	return d.FindOne(findOption)
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

	tp, err = tpTbl.ToModel()
	if err != nil {
		return
	}

	return
}
