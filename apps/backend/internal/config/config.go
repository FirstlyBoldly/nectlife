package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/nectgrams-webapp-team/nectlife/internal/must"
)

type config struct {
	POSTGRES_DB_URL                       string
	SESSION_COOKIE_NAME                   string
	GARBAGE_COLLECTOR_INTERVAL_IN_MINUTES int
	EXPIRES_AT_IN_HOURS                   int
	ALLOWED_ORIGINS                       []string
	HOST_PORT                             int
}

func Load() *config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	psqlDbURL := must.LookupEnvEnforce("POSTGRES_DB_URL")
	sessCookieName := must.LookupEnvEnforce("SESSION_COOKIE_NAME")
	gbInterval := must.AtoiEnforce(must.LookupEnvEnforce("GARBAGE_COLLECTOR_INTERVAL_IN_MINUTES"))
	expiresAt := must.AtoiEnforce(must.LookupEnvEnforce("EXPIRES_AT_IN_HOURS"))
	allowedOrigins := strings.Split(must.LookupEnvEnforce("ALLOWED_ORIGINS"), ",")
	hostP := must.AtoiEnforce(must.LookupEnvEnforce("BACKEND_HOST_PORT"))

	return &config{
		POSTGRES_DB_URL:                       psqlDbURL,
		SESSION_COOKIE_NAME:                   sessCookieName,
		GARBAGE_COLLECTOR_INTERVAL_IN_MINUTES: gbInterval,
		EXPIRES_AT_IN_HOURS:                   expiresAt,
		ALLOWED_ORIGINS:                       allowedOrigins,
		HOST_PORT:                             hostP,
	}
}

var Data *config = Load()
