package session

import (
	"context"
	"net/http"
	"time"

	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type NlHandler func(next http.Handler) http.Handler

type SessionHandlerInterface interface {
	Start(*http.Request) (*postgres.Session, *http.Request)
	ValidateSession(context.Context, *postgres.Session) bool
	Migrate(context.Context, *postgres.Session) error
	GarbageCollect(time.Duration)
	Handle(http.Handler) http.Handler
}
