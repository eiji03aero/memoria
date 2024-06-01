package model

import "time"

type Thread struct {
	ID          string
	UserSpaceID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewThreadDTO struct {
	ID          string
	UserSpaceID string
}

func NewThread(dto NewThreadDTO) *Thread {
	return &Thread{
		ID:          dto.ID,
		UserSpaceID: dto.UserSpaceID,
	}
}
