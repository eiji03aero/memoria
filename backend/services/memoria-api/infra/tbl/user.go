package tbl

import (
	"time"

	"memoria-api/domain/model"
)

type User struct {
	ID           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Email        string    `gorm:"column:email"`
	PasswordHash string    `gorm:"column:password_hash"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (t User) TableName() string {
	return "users"
}

func (t User) ToModel() *model.User {
	return model.NewUser(model.NewUserDTO{
		ID:           t.ID,
		Name:         t.Name,
		Email:        t.Email,
		PasswordHash: t.PasswordHash,
	})
}
