package repository

import "memoria-api/domain/model"

type Timeline interface {
	Find(fOpt *FindOption) (tus []*model.TimelineUnit, err error)
}
