package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type UserSpaceAlbumRelation struct {
	UserSpaceID string    `gorm:"column:user_space_id"`
	AlbumID     string    `gorm:"column:album_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t UserSpaceAlbumRelation) TableName() string {
	return "user_space_album_relations"
}

func (t UserSpaceAlbumRelation) ToModel() *model.UserSpaceAlbumRelation {
	return model.NewUserSpaceAlbumRelation(model.NewUserSpaceAlbumRelationDTO{
		UserSpaceID: t.UserSpaceID,
		AlbumID:     t.AlbumID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	})
}
