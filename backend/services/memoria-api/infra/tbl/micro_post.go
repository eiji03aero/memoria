package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type MicroPost struct {
	ID          string    `gorm:"column:id"`
	UserID      string    `gorm:"column:user_id"`
	UserSpaceID string    `gorm:"column:user_space_id"`
	Content     string    `gorm:"column:content"`
	Media       []*Medium `gorm:"many2many:micro_post_medium_relations;foreignKey:ID;joinForeignKey:micro_post_id;References:ID;joinReferences:medium_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (t MicroPost) TableName() string {
	return "micro_posts"
}

func (t MicroPost) ToModel() (m *model.MicroPost) {
	m = model.NewMicroPost(model.NewMicroPostDTO{
		ID:          t.ID,
		UserID:      t.UserID,
		UserSpaceID: t.UserSpaceID,
		Content:     t.Content,
	})

	for _, mediumTbl := range t.Media {
		m.Media = append(m.Media, mediumTbl.ToModel())
	}

	m.CreatedAt = t.CreatedAt
	m.UpdatedAt = t.UpdatedAt
	return
}

func (t *MicroPost) FromModel(m *model.MicroPost) {
	t.ID = m.ID
	t.UserID = m.UserID
	t.UserSpaceID = m.UserSpaceID
	t.Content = m.Content
	t.CreatedAt = m.CreatedAt
	t.UpdatedAt = m.UpdatedAt
}
