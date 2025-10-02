package routes

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nectgrams-webapp-team/nectlife/internal/auth"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	nlerrs "github.com/nectgrams-webapp-team/nectlife/internal/errors"
	nlms "github.com/nectgrams-webapp-team/nectlife/internal/middlewares"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
	"github.com/nectgrams-webapp-team/nectlife/internal/session"
	"github.com/nectgrams-webapp-team/nectlife/internal/user"
)

func NewRouter(db *postgres.Queries) (*chi.Mux, huma.API) {
	sessSvc := session.NewSessionStore(db)

	interval := time.Duration(config.Data.GARBAGE_COLLECTOR_INTERVAL_IN_MINUTES) * time.Minute
	sessHndlr := session.NewSessionHandler(interval, sessSvc, config.Data.SESSION_COOKIE_NAME)

	authHndlr := auth.NewAuthHandler(db, sessSvc, sessHndlr)
	userHndlr := user.NewUserHandler(db)

	c := newCors()

	mux := chi.NewRouter()
	mux.Use(c.Handler)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(nlms.URLCleaner)
	mux.Use(sessHndlr.Handle)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(30 * time.Second))

	api := newHumaApi(mux, config.Data.SESSION_COOKIE_NAME)
	userHndlr.Routes(mux, api, authHndlr)
	authHndlr.Routes(mux, api)
	mux.NotFound(nlerrs.NotFoundHandler)

	return mux, api
}
