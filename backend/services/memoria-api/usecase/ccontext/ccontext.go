package ccontext

import (
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"

	"github.com/gin-gonic/gin"
)

var (
	KeyUserID      = "cctx-user-id"
	KeyUserSpaceID = "cctx-user-space-id"
)

type context struct {
	userID      string
	userSpaceID string
}

func NewContext(ctx *gin.Context) interfaces.Context {
	return &context{
		userID:      GetUserID(ctx),
		userSpaceID: GetUserSpaceID(ctx),
	}
}

func (c *context) GetUserID() string {
	if c.userID == "" {
		panic(cerrors.NewInternal("user id is not set").Error())
	}
	return c.userID
}

func GetUserID(ctx *gin.Context) string {
	userID, ok := ctx.MustGet(KeyUserID).(string)
	if !ok {
		panic(cerrors.NewInternal("user id is not set").Error())
	}
	return userID
}

func SetUserID(ctx *gin.Context, userID string) {
	ctx.Set(KeyUserID, userID)
}

func (c *context) GetUserSpaceID() string {
	if c.userSpaceID == "" {
		panic(cerrors.NewInternal("user spaceid is not set").Error())
	}
	return c.userSpaceID
}

func GetUserSpaceID(ctx *gin.Context) string {
	userSpaceID, ok := ctx.MustGet(KeyUserSpaceID).(string)
	if !ok {
		panic(cerrors.NewInternal("user space id is not set").Error())
	}
	return userSpaceID
}

func SetUserSpaceID(ctx *gin.Context, userSpaceID string) {
	ctx.Set(KeyUserSpaceID, userSpaceID)
}
