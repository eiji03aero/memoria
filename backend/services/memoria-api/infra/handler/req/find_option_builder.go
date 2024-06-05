package req

import (
	"memoria-api/domain/interfaces/repository"

	"github.com/gin-gonic/gin"
)

func BuildFindOptionWithQuery(c *gin.Context, findOpt *repository.FindOption) (err error) {
	pagiQ := &Paginate{}
	err = c.ShouldBindQuery(pagiQ)
	if err != nil {
		return
	}

	if pagiQ.Page != nil && pagiQ.PerPage != nil {
		offset := (*pagiQ.Page - 1) * *pagiQ.PerPage
		findOpt.Offset = &offset
		findOpt.Limit = pagiQ.PerPage
	}

	cpagiQ := &CPaginate{}
	err = c.ShouldBindQuery(cpagiQ)
	if err != nil {
		return
	}

	if cpagiQ.Cursor != nil {
		if cpagiQ.CBefore != nil {
			findOpt.Cursor = *cpagiQ.Cursor
			findOpt.CBefore = *cpagiQ.CBefore
		}
		if cpagiQ.CAfter != nil {
			findOpt.Cursor = *cpagiQ.Cursor
			findOpt.CAfter = *cpagiQ.CAfter
		}

		if cpagiQ.CExclude != nil {
			findOpt.CExclude = *cpagiQ.CExclude == "true"
		}
	}

	return
}
