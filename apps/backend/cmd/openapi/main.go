package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"github.com/nectgrams-webapp-team/nectlife/internal/routes"
	"gopkg.in/yaml.v3"
)

func main() {
	slog.Info("Generating OpenAPI docs...")
	pool, err := pgxpool.New(context.Background(), config.Data.POSTGRES_DB_URL)
	if err != nil {
		panic(err)
	}

	defer pool.Close()
	q := postgres.New(pool)
	_, spec := routes.NewRouter(q)

	yamlBytes, err := yaml.Marshal(spec.OpenAPI())
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("api/openapi.yaml", yamlBytes, 0644); err != nil {
		panic(err)
	}

	slog.Info("Generated OpenAPI docs!")
}
