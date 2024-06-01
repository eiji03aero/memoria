package model

import "time"

type MicroPost struct {
	ID          string
	UserID      string
	UserSpaceID string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewMicroPostDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Content     string
}

func NewMicroPost(dto NewMicroPostDTO) *MicroPost {
	return &MicroPost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Content:     dto.Content,
	}
}
