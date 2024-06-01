package repository

import "memoria-api/domain/model"

type ThreadMicroPostRelation interface {
	Create(dto ThreadMicroPostRelationDTO) (tmpr *model.ThreadMicroPostRelation, err error)
}

type ThreadMicroPostRelationDTO struct {
	ThreadID    string
	MicroPostID string
}
