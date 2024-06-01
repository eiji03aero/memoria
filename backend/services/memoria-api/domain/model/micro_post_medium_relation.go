package model

import "time"

type MicroPostMediumRelation struct {
	MicroPostID string
	MediumID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewMicroPostMediumRelationDTO struct {
	MicroPostID string
	MediumID    string
}

func NewMicroPostMediumRelation(dto NewMicroPostMediumRelationDTO) *MicroPostMediumRelation {
	return &MicroPostMediumRelation{
		MicroPostID: dto.MicroPostID,
		MediumID:    dto.MediumID,
	}
}
