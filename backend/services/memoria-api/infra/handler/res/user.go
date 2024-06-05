package res

import "memoria-api/domain/model"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserFromModel(m *model.User) *User {
	return &User{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}
