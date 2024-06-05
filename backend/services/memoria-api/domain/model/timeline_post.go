package model

import "time"

type TimelinePost struct {
	ID          string
	UserID      string
	UserSpaceID string
	User        *User
	Thread      *Thread
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewTimelinePostDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	User        *User
	Thread      *Thread
}

func NewTimelinePost(dto NewTimelinePostDTO) *TimelinePost {
	return &TimelinePost{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		User:        dto.User,
		Thread:      dto.Thread,
	}
}
