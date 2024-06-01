package repository

import "memoria-api/domain/model"

type TimelinePostThreadRelation interface {
	Find(fOpt *FindOption) (tptrs []*model.TimelinePostThreadRelation, err error)
	Create(dto TimelinePostThreadRelationCreateDTO) (tptr *model.TimelinePostThreadRelation, err error)
}

type TimelinePostThreadRelationCreateDTO struct {
	TimelinePostID string
	ThreadID       string
}
