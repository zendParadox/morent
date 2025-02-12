package config

import (
	"os"
)

var SecretKey = os.Getenv("JWT_SECRET")
