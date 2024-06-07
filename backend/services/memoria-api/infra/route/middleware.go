package route

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"memoria-api/application/ccontext"
	"memoria-api/application/usecase"
	"memoria-api/domain/cerrors"
	"memoria-api/domain/interfaces"
	"memoria-api/infra/handler/res"
	"memoria-api/infra/registry"
	"memoria-api/util"
)

func wrap(regb *registry.Builder, h func(c *gin.Context, reg interfaces.Registry) (status int, data any, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		reg, err := regb.Build(registry.BuilderBuildDTO{InitDB: util.BoolToPointer(true)})
		if err != nil {
			panic(err)
		}

		lgr := reg.NewLogger()
		lgr.Info(fmt.Sprintf("API Starting: %s %s", c.Request.Method, c.Request.URL.Path))
		defer func() {
			lgr.Info("API finished")
			reg.CloseDB()
		}()

		// -------------------- handler --------------------
		status, data, err := h(c, reg)
		if err != nil {
			lgr.Error("wrap handler result err: ", err.Error())

			if errors.As(err, &cerrors.Validation{}) {
				c.JSON(http.StatusBadRequest, res.NewValidationRes(err.(cerrors.Validation)))
				c.Abort()
				return
			}

			if errors.As(err, &cerrors.Internal{}) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": cerrors.NewInternal(err.Error()).Error(),
			})
			c.Abort()
			return
		}

		c.JSON(status, data)
		return
	}
}

func Authenticate(regb *registry.Builder) gin.HandlerFunc {
	return func(c *gin.Context) {
		reg, err := regb.Build(registry.BuilderBuildDTO{InitDB: util.BoolToPointer(true)})
		defer reg.CloseDB()

		authUc := usecase.NewAuth(reg)

		tokenString := c.GetHeader("Authorization")
		leadingLen := len("Bearer ")
		if tokenString == "" || len(tokenString) <= leadingLen {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": cerrors.NewUnauthorized().Error(),
			})
			c.Abort()
			return
		}

		tokenString = tokenString[leadingLen:]

		ret, err := authUc.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": cerrors.NewUnauthorized().Error(),
			})
			c.Abort()
			return
		}

		ccontext.SetUserID(c, ret.UserID)
		ccontext.SetUserSpaceID(c, ret.UserSpaceID)

		c.Next()
	}
}
