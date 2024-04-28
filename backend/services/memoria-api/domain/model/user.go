package model

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
}

type NewUserDTO struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
}

func NewUser(dto NewUserDTO) *User {
	return &User{
		ID:           dto.ID,
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: dto.PasswordHash,
	}
}
