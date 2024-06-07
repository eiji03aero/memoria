package res

import "memoria-api/domain/model"

type Album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AlbumFromModel(m *model.Album) (r *Album) {
	r = &Album{}
	r.ID = m.ID
	r.Name = m.Name
	return
}
