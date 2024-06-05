package repository

import "memoria-api/domain/model"

type Timeline interface {
	Find(fOpt *FindOption) (tus model.TimelineUnits, cpagi CPagination, err error)
}
