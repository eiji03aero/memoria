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
	Units []*res.TimelineUnit `json:"units"`
	CPagi res.CPagination     `json:"cpagination"`
}

func (h *Timeline) Find(c *gin.Context, reg interfaces.Registry) (status int, data any, err error) {
	cctx := ccontext.NewContext(c)

	fOpt := &repository.FindOption{
		Filter: map[string]any{
			"user_space_id": cctx.GetUserSpaceID(),
		},
	}
	err = req.BuildFindOptionWithQuery(c, fOpt)
	if err != nil {
		return
	}

	units, cpagi, err := reg.NewTimelineRepository().Find(fOpt)
	if err != nil {
		return
	}

	timelineFindRes := TimelineFindRes{
		CPagi: res.CPagination{
			PrevCursor: cpagi.PrevCursor,
			NextCursor: cpagi.NextCursor,
		},
		Units: []*res.TimelineUnit{},
	}
	for _, tu := range units {
		timelineFindRes.Units = append(timelineFindRes.Units, res.TimelineUnitFromModel(tu))
	}

	status = http.StatusOK
	data = timelineFindRes
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
