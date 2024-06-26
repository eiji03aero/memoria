package svc

type Album interface {
	AddMedia(dto AlbumAddMediaDTO) error
	RemoveMedia(dto AlbumRemoveMediaDTO) error
}

type AlbumAddMediaDTO struct {
	UserSpaceID string
	AlbumID     string
	MediumIDs   []string
}

type AlbumRemoveMediaDTO struct {
	AlbumID   string
	MediumIDs []string
}
