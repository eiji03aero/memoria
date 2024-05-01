package config

import (
	"os"
	"strings"
)

var JWTSecretKey = os.Getenv("JWT_SECRET_KEY")

var CORSAllowOrigins = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
