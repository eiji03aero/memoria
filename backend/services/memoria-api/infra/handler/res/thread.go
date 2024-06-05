package res

import (
	"memoria-api/domain/model"
)

type Thread struct {
	ID         string       `json:"id"`
	MicroPosts []*MicroPost `json:"micro_posts"`
}

func ThreadFromModel(m *model.Thread) (t *Thread) {
	t = &Thread{
		ID: m.ID,
	}

	for _, mp := range m.MicroPosts {
		t.MicroPosts = append(t.MicroPosts, MicroPostFromModel(mp))
	}

	return
}
