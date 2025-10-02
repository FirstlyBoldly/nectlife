package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmittmann/tint"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"github.com/nectgrams-webapp-team/nectlife/internal/routes"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			AddSource:  true,
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
		}),
	))

	pool, err := pgxpool.New(context.Background(), config.Data.POSTGRES_DB_URL)
	if err != nil {
		panic(err)
	}

	defer pool.Close()
	db := postgres.New(pool)

	router, _ := routes.NewRouter(db)
	port := fmt.Sprintf(":%d", config.Data.HOST_PORT)

	slog.Info("Listening on port: " + port)
	http.ListenAndServe(port, router)
}
