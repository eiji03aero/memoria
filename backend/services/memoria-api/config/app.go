package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// -------------------- For local env --------------------
	if _, err := os.Stat(".env.local"); errors.Is(err, os.ErrNotExist) {
		// If file not found, just early exit
		return
	}

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local")
	}
}

var (
	Env        = os.Getenv("ENV")
	Host       = os.Getenv("HOST")
	ClientHost = os.Getenv("CLIENT_HOST")

	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
	DBLogLevel = os.Getenv("DB_LOG_LEVEL")

	JWTSecretKey     = os.Getenv("JWT_SECRET_KEY")
	CORSAllowOrigins = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")

	NoReplyEmailAddress = "no-reply@memoria-app.com"
)
