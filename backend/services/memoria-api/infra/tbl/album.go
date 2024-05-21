package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type Album struct {
	ID        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (t Album) TableName() string {
	return "albums"
}

func (t Album) ToModel() (album *model.Album) {
	album = model.NewAlbum(model.NewAlbumDTO{
		ID:   t.ID,
		Name: t.Name,
	})

	album.CreatedAt = t.CreatedAt
	album.UpdatedAt = t.UpdatedAt
	return
}

func (t *Album) FromModel(album *model.Album) {
	t.ID = album.ID
	t.Name = album.Name
	t.CreatedAt = album.CreatedAt
	t.UpdatedAt = album.UpdatedAt
}
