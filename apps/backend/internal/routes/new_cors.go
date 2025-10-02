package routes

import (
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/rs/cors"
)

func newCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   config.Data.ALLOWED_ORIGINS,
		AllowedMethods:   []string{"GET", "POST", "UPDATE", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Accept", "Cookie"},
		AllowCredentials: true,
		Debug:            false,
	})
}
