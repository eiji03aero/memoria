package model

import (
	"time"

	"memoria-api/domain/value"
)

type User struct {
	ID            string
	AccountStatus value.UserAccountStatus
	Name          string
	Email         string
	PasswordHash  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type NewUserDTO struct {
	ID            string
	AccountStatus value.UserAccountStatus
	Name          string
	Email         string
	PasswordHash  string
}

func NewUser(dto NewUserDTO) (*User, error) {
	return &User{
		ID:            dto.ID,
		AccountStatus: dto.AccountStatus,
		Name:          dto.Name,
		Email:         dto.Email,
		PasswordHash:  dto.PasswordHash,
	}, nil
}

func (m *User) SetAccountStatus(status value.UserAccountStatus) {
	m.AccountStatus = status
}

func (user *User) IsStatusValidForUse() bool {
	return user.AccountStatus == value.UserAccountStatus_Confirmed
}
