package model

import "time"

type ThreadMicroPostRelation struct {
	ThreadID    string
	MicroPostID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewThreadMicroPostRelationDTO struct {
	ThreadID    string
	MicroPostID string
}

func NewThreadMicroPostRelation(dto NewThreadMicroPostRelationDTO) *ThreadMicroPostRelation {
	return &ThreadMicroPostRelation{
		ThreadID:    dto.ThreadID,
		MicroPostID: dto.MicroPostID,
	}
}
