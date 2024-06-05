package model

import "time"

type MicroPost struct {
	ID          string
	UserID      string
	UserSpaceID string
	Content     string
	Media       []*Medium
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewMicroPostDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Content     string
	Media       []*Medium
}

func NewMicroPost(dto NewMicroPostDTO) *MicroPost {
	return &MicroPost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
		Media:       dto.Media,
	}
}
