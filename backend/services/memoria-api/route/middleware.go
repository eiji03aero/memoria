package route

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"memoria-api/domain/cerrors"
	"memoria-api/registry"
	"memoria-api/usecase"
	"memoria-api/usecase/ccontext"
)

func wrap(h func(c *gin.Context, reg registry.Registry) (status int, data any, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}

		reg := registry.NewRegistry(db)
		status, data, err := h(c, reg)
		if err != nil {
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
		authUc := usecase.NewAuth()

		tokenString := c.GetHeader("Authorization")
		leadingLen := len("Bearer ")
		log.Println("tokenString", tokenString)
		if tokenString == "" || len(tokenString) <= leadingLen {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": cerrors.NewUnauthorized().Error(),
			})
			c.Abort()
			return
		}

		tokenString = tokenString[leadingLen:]

		userID, userSpaceID, err := authUc.VerifyJWT(tokenString)
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
