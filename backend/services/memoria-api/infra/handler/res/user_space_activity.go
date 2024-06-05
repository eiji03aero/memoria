package res

import (
	"encoding/json"
	"time"

	"memoria-api/domain/model"
)

type UserSpaceActivity struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Data      any       `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func UserSpaceActivityFromModel(m *model.UserSpaceActivity) (usa *UserSpaceActivity, err error) {
	data := map[string]any{}
	err = json.Unmarshal([]byte(m.Data), &data)
	if err != nil {
		return
	}

	usa = &UserSpaceActivity{
		ID:        m.ID,
		Type:      string(m.Type),
		Data:      data,
		CreatedAt: m.CreatedAt,
	}
	return
}
