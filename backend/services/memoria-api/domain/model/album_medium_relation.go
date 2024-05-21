package model

import "time"

type AlbumMediumRelation struct {
	AlbumID   string
	MediumID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewAlbumMediumRelationDTO struct {
	AlbumID  string
	MediumID string
}

func NewAlbumMediumRelation(dto NewAlbumMediumRelationDTO) *AlbumMediumRelation {
	return &AlbumMediumRelation{
		AlbumID:  dto.AlbumID,
		MediumID: dto.MediumID,
	}
}
