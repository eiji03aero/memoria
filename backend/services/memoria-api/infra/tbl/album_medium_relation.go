package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type AlbumMediumRelation struct {
	AlbumID   string    `gorm:"column:album_id"`
	MediumID  string    `gorm:"column:medium_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (t AlbumMediumRelation) TableName() string {
	return "album_medium_relations"
}

func (t AlbumMediumRelation) ToModel() (amr *model.AlbumMediumRelation) {
	amr = model.NewAlbumMediumRelation(model.NewAlbumMediumRelationDTO{
		AlbumID:  t.AlbumID,
		MediumID: t.MediumID,
	})
	amr.CreatedAt = t.CreatedAt
	amr.UpdatedAt = t.UpdatedAt
	return
}
