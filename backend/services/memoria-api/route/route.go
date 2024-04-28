package route

import (
	"memoria-api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	// -------------------- Public apis --------------------
	public := api.Group("/public")

	{
		account := handler.NewAccount()
		public.POST("/signup", wrap(account.Signup))
	}

	// -------------------- Authenticated apis --------------------
	authenticated := api.Group("/auth", Authenticate())

	{
		appData := handler.NewAppData()
		authenticated.GET("/app-data", wrap(appData.Get))
	}

	return router
}
