package route

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"memoria-api/domain/cerrors"
	"memoria-api/infra/caws"
	"memoria-api/infra/db"
	"memoria-api/registry"
	"memoria-api/usecase"
	"memoria-api/usecase/ccontext"
)

func buildRegistry(ctx context.Context) (reg registry.Registry, err error) {
	// -------------------- database --------------------
	db, err := db.New()
	if err != nil {
		return
	}

	// -------------------- aws --------------------
	awsCfg, err := caws.LoadConfig(ctx)
	if err != nil {
		return
	}

	// -------------------- registry --------------------
	reg = registry.NewRegistry(registry.NewRegistryDTO{
		DB:     db,
		AWSCfg: awsCfg,
	})

	return
}

func wrap(h func(c *gin.Context, reg registry.Registry) (status int, data any, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(fmt.Sprintf(
			"API Starting: %s %s",
			c.Request.Method, c.Request.URL.Path,
		))
		defer func() {
			log.Println("API finished")
		}()

		ctx := context.Background()
		reg, err := buildRegistry(ctx)
		if err != nil {
			log.Println("wrap build registry error:", err.Error())
			return
		}

		// -------------------- handler --------------------
		status, data, err := h(c, reg)
		if err != nil {
			log.Println("wrap handler result err: ", err.Error())

			if errors.Is(err, cerrors.Validation{}) {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": cerrors.NewValidation(err.Error()).Error(),
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

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		reg, err := buildRegistry(ctx)
		if err != nil {
			log.Println("Authenticate build registry error:", err.Error())
			return
		}

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

		userID, userSpaceID, err := authUc.VerifyJWT(usecase.AuthVerifyJWTDTO{
			TokenString: tokenString,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": cerrors.NewUnauthorized().Error(),
			})
			c.Abort()
			return
		}

		ccontext.SetUserID(c, userID)
		ccontext.SetUserSpaceID(c, userSpaceID)

		c.Next()
	}
}
