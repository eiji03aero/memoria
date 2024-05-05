package route

import (
	"net/http"

	"memoria-api/config"
	"memoria-api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORSAllowOrigins,
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	// -------------------- Public apis --------------------
	public := api.Group("/public")

	public.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	{
		account := handler.NewAccount()
		public.POST("/signup", wrap(account.Signup))
		public.GET("/signup-confirm", wrap(account.SignupConfirm))
		public.POST("/login", wrap(account.Login))
		public.POST("/invite-user-confirm", wrap(account.InviteUserConfirm))
	}

	// -------------------- Authenticated apis --------------------
	authenticated := api.Group("/auth", Authenticate())

	{
		appData := handler.NewAppData()
		authenticated.GET("/app-data", wrap(appData.Get))

		account := handler.NewAccount()
		authenticated.POST("/invite-user", wrap(account.InviteUser))
	}

	return router
}
