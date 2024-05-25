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

	return
}
