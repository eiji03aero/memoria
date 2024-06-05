package repository

import "memoria-api/domain/model"

type TimelinePost interface {
	Find(fOpt *FindOption) (tps []*model.TimelinePost, err error)
	FindOne(fOpt *FindOption) (tp *model.TimelinePost, err error)
	FindOneByID(id string, fOpt *FindOption) (tp *model.TimelinePost, err error)
	Create(dto TimelinePostCreateDTO) (tp *model.TimelinePost, err error)
}

type TimelinePostCreateDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
}
