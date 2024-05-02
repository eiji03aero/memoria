package tbl

import (
	"time"

	"memoria-api/domain/model"
	"memoria-api/domain/value"
)

type User struct {
	ID            string    `gorm:"column:id"`
	AccountStatus string    `gorm:"column:account_status"`
	Name          string    `gorm:"column:name"`
	Email         string    `gorm:"column:email"`
	PasswordHash  string    `gorm:"column:password_hash"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (t User) TableName() string {
	return "users"
}

func (t User) ToModel() (user *model.User, err error) {
	accountStatus, err := value.NewUserAccountStatus(t.AccountStatus)
	if err != nil {
		return
	}

	user, err = model.NewUser(model.NewUserDTO{
		ID:            t.ID,
		AccountStatus: accountStatus,
		Name:          t.Name,
		Email:         t.Email,
		PasswordHash:  t.PasswordHash,
	})
	if err != nil {
		return
	}

	user.CreatedAt = t.CreatedAt
	user.UpdatedAt = t.UpdatedAt
	return
}

func (t *User) FromModel(user *model.User) {
	t.ID = user.ID
	t.AccountStatus = string(user.AccountStatus)
	t.Name = user.Name
	t.Email = user.Email
	t.PasswordHash = user.PasswordHash
	t.CreatedAt = user.CreatedAt
	t.UpdatedAt = user.UpdatedAt
}
