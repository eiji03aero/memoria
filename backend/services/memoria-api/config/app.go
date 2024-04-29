package config

import (
	"os"
	"strings"
)

var JWTSecretKey = "HogeZamurai"

var CORSAllowOrigins = strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
