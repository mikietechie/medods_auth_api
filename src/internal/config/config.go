/*
Date Created		28 Nov 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Status				Looking for a job!
Description			A Golang Auth API
*/

package config

import (
	"context"
	"time"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload()
var CTX = context.Background()

/* System */
var SYS_NAME = GetEnvOrDef("SYS_NAME", "Go Auth API")
var SECRET_KEY = GetEnvOrDef("SECRET_KEY", "$UPER_$EXRE8!")
var ENV = GetEnvOrDef("ENV", "PROD")
var DEV = ENV == "DEV"
var SERVER_ADDRESS = GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")

var DATABASE_CONNECTION = GetEnvOrDef(
	"POSTGRES_CONNECTION",
	"postgres://pg:pass@localhost:5432/db",
)
var JWT_ACCESS_TOKEN_LIFETIME = time.Minute * 5
var JWT_REFRESH_TOKEN_LIFETIME = time.Hour * 24 * 7
