package model

import "time"

type TimelinePost struct {
	ID          string
	UserID      string
	UserSpaceID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewTimelinePostDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
}

func NewTimelinePost(dto NewTimelinePostDTO) *TimelinePost {
	return &TimelinePost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
	}
}
