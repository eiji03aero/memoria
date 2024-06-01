package repository

import "memoria-api/domain/model"

type MicroPostMediumRelation interface {
	Find(fOpt *FindOption) (mpmrs []*model.MicroPostMediumRelation, err error)
	Create(dto MicroPostMediumRelationCreateDTO) (mpmr *model.MicroPostMediumRelation, err error)
}

type MicroPostMediumRelationCreateDTO struct {
	MicroPostID string
	MediumID    string
}
