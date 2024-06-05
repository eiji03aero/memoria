package res

import (
	"memoria-api/domain/model"
)

type MicroPost struct {
	ID      string    `json:"id"`
	Content string    `json:"content"`
	Media   []*Medium `json:"media"`
}

func MicroPostFromModel(m *model.MicroPost) (mp *MicroPost) {
	mp = &MicroPost{
		ID:      m.ID,
		Content: m.Content,
	}

	for _, medium := range m.Media {
		mp.Media = append(mp.Media, MediumFromModel(medium))
	}
	return
}
