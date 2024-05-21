package handler

import (
	"net/http"

	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type AppData struct{}

func NewAppData() *AppData {
	return &AppData{}
}

type AppDataGetResUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type AppDataGetResUserSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type AppDataGetRes struct {
	User      AppDataGetResUser      `json:"user"`
	UserSpace AppDataGetResUserSpace `json:"user_space"`
}

func (h *AppData) Get(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)

	appDataUc := usecase.NewAppData(reg)
	aggregated, err := appDataUc.Get(usecase.AppDataGetDTO{
		UserID:      cctx.GetUserID(),
		UserSpaceID: cctx.GetUserSpaceID(),
	})
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	data = AppDataGetRes{
		User: AppDataGetResUser{
			ID:   aggregated.User.ID,
			Name: aggregated.User.Name,
		},
		UserSpace: AppDataGetResUserSpace{
			ID:   aggregated.UserSpace.ID,
			Name: aggregated.UserSpace.Name,
		},
	}

	return
}
