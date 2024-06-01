package handler

import (
	"net/http"

	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/interfaces"
	"memoria-api/domain/interfaces/repository"
	"memoria-api/infra/handler/req"
	"memoria-api/infra/handler/res"

	"github.com/gin-gonic/gin"
)

type Timeline struct{}

func NewTimeline() *Timeline {
	return &Timeline{}
}

type TimelineFindRes struct {
	Pagi  res.Pagination
	Units []*res.TimelineUnit `json:"units"`
}

func (h *Timeline) Find(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	fOpt := &repository.FindOption{}
	err = req.BuildFindOptionWithQuery(c, fOpt)
	if err != nil {
		return
	}

	_, err = reg.NewTimelineRepository().Find(fOpt)
	if err != nil {
		return
	}

	status = http.StatusOK
	data = TimelineFindRes{}
	return
}

type TimelineCreatePostReq struct {
	Content   string   `json:"content"`
	MediumIDs []string `json:"medium_ids"`
}

func (h *Timeline) CreatePost(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)
	timelineUc := usecase.NewTimeline(reg)

	body := TimelineCreatePostReq{}
	err = c.BindJSON(&body)
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	_, err = timelineUc.CreatePost(usecase.TimelineCreatePostDTO{
		UserID:      cctx.GetUserID(),
		UserSpaceID: cctx.GetUserSpaceID(),
		Content:     body.Content,
		MediumIDs:   body.MediumIDs,
	})
	if err != nil {
		return
	}

	status = http.StatusOK
	return
}
