package model

import "time"

type Album struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewAlbumDTO struct {
	ID   string
	Name string
}

func NewAlbum(dto NewAlbumDTO) *Album {
	return &Album{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
