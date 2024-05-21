package res

import "memoria-api/domain/model"

type Album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (a *Album) FromModel(album *model.Album) *Album {
	a.ID = album.ID
	a.Name = album.Name
	return a
}
