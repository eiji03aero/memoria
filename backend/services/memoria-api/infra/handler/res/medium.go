package res

import "memoria-api/domain/model"

type Medium struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Extension   string `json:"extension"`
	OriginalURL string `json:"original_url"`
	Tn240URL    string `json:"tn_240_url"`
}

func MediumFromModel(m *model.Medium) *Medium {
	return &Medium{
		ID:          m.ID,
		Name:        m.Name,
		Extension:   m.Extension,
		OriginalURL: m.GetOriginalURL(),
		Tn240URL:    m.GetTn240URL(),
	}
}

type UploadURL struct {
	MediumID string `json:"medium_id"`
	URL      string `json:"url"`
}
