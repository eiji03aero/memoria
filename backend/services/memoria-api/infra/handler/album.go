package handler

import (
	"net/http"

	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/handler/res"

	"github.com/gin-gonic/gin"
)

type Album struct{}

func NewAlbum() *Album {
	return &Album{}
}

type AlbumFindRes struct {
	Albums []*res.Album `json:"albums"`
}

func (h *Album) Find(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)

	albums, err := reg.NewAlbumRepository().Find(&repository.FindOption{
		Joins: []*repository.FindOptionJoin{
			{Query: "join user_space_album_relations usar on usar.album_id = albums.id"},
		},
		Filters: []*repository.FindOptionFilter{
			{Query: "usar.user_space_id = ?", Value: cctx.GetUserSpaceID()},
		},
		Order: "id",
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	albumRess := make([]*res.Album, 0, len(albums))
	for _, album := range albums {
		albumRess = append(albumRess, res.AlbumFromModel(album))
	}
	data = AlbumFindRes{
		Albums: albumRess,
	}
	return
}

type AlbumFindOneRes struct {
	Album *res.Album `json:"album"`
}

func (h *Album) FindOne(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)

	albumID := c.Param("id")

	album, err := reg.NewAlbumRepository().FindOne(&repository.FindOption{
		Filters: []*repository.FindOptionFilter{
			{Query: "usar.user_space_id = ?", Value: cctx.GetUserSpaceID()},
			{Query: "id = ?", Value: albumID},
		},
		Joins: []*repository.FindOptionJoin{
			{Query: "join user_space_album_relations usar on usar.album_id = albums.id"},
		},
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	data = AlbumFindOneRes{
		Album: res.AlbumFromModel(album),
	}
	return
}

type AlbumCreateReq struct {
	Name *string `json:"name"`
}

type AlbumCreateRes struct {
	Album *res.Album `json:"album"`
}

func (h *Album) Create(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	albumUc := usecase.NewAlbum(reg)

	body := AlbumCreateReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	createRet, err := albumUc.Create(usecase.AlbumCreateDTO{
		Name:        body.Name,
		UserSpaceID: cctx.GetUserSpaceID(),
	})
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	status = http.StatusOK
	data = AlbumCreateRes{
		Album: res.AlbumFromModel(createRet.Album),
	}
	return
}

func (h *Album) Delete(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	albumUc := usecase.NewAlbum(reg)

	id := c.Param("id")

	err = albumUc.Delete(usecase.AlbumDeleteDTO{
		ID: id,
	})
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	status = http.StatusOK
	return
}

type AlbumAddMediaReq struct {
	MediumIDs []string `json:"medium_ids"`
	AlbumIDs  []string `json:"album_ids"`
}

func (h *Album) AddMedia(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	albumUc := usecase.NewAlbum(reg)

	body := AlbumAddMediaReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	err = albumUc.AddMedia(usecase.AlbumAddMediaDTO{
		UserSpaceID: cctx.GetUserSpaceID(),
		AlbumIDs:    body.AlbumIDs,
		MediumIDs:   body.MediumIDs,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	return
}

type AlbumRemoveMediaReq struct {
	MediumIDs []string `json:"medium_ids"`
	AlbumIDs  []string `json:"album_ids"`
}

func (h *Album) RemoveMedia(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	albumUc := usecase.NewAlbum(reg)

	body := AlbumAddMediaReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	err = albumUc.RemoveMedia(usecase.AlbumRemoveMediaDTO{
		AlbumIDs:  body.AlbumIDs,
		MediumIDs: body.MediumIDs,
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	return
}
