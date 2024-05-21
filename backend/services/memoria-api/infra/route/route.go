package route

import (
	"net/http"

	"memoria-api/config"
	"memoria-api/infra/handler"
	"memoria-api/infra/registry"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(regb *registry.Builder) *gin.Engine {
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
		public.POST("/signup", wrap(regb, account.Signup))
		public.GET("/signup-confirm", wrap(regb, account.SignupConfirm))
		public.POST("/login", wrap(regb, account.Login))
		public.POST("/invite-user-confirm", wrap(regb, account.InviteUserConfirm))
	}

	// -------------------- Authenticated apis --------------------
	authenticated := api.Group("/auth", Authenticate(regb))

	{
		// -------------------- app data --------------------
		appData := handler.NewAppData()
		authenticated.GET("/app-data", wrap(regb, appData.Get))

		// -------------------- account --------------------
		account := handler.NewAccount()
		authenticated.POST("/invite-user", wrap(regb, account.InviteUser))

		// -------------------- album --------------------
		album := handler.NewAlbum()
		authenticated.GET("/albums/:id", wrap(regb, album.FindOne))
		authenticated.GET("/albums", wrap(regb, album.Find))
		authenticated.POST("/albums", wrap(regb, album.Create))
		authenticated.DELETE("/albums/:id", wrap(regb, album.Delete))
		authenticated.POST("/albums/add-media", wrap(regb, album.AddMedia))
		authenticated.POST("/albums/remove-media", wrap(regb, album.RemoveMedia))

		// -------------------- medium --------------------
		medium := handler.NewMedium()
		authenticated.GET("/media", wrap(regb, medium.Find))
		authenticated.POST("/media/request-upload-urls", wrap(regb, medium.RequestUploadURLs))
		authenticated.POST("/media/confirm-uploads", wrap(regb, medium.ConfirmUploads))
		authenticated.DELETE("/media/:id", wrap(regb, medium.Delete))
	}

	return router
}
