package model

import "time"

type Thread struct {
	ID          string
	UserSpaceID string
	MicroPosts  []*MicroPost
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewThreadDTO struct {
	ID          string
	UserSpaceID string
	MicroPosts  []*MicroPost
}

func NewThread(dto NewThreadDTO) *Thread {
	return &Thread{
		ID:          dto.ID,
		UserSpaceID: dto.UserSpaceID,
		MicroPosts:  dto.MicroPosts,
	}
}
