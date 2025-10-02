package session

import (
	"context"

	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type SessionServiceInterface interface {
	CreateSession(ctx context.Context) (postgres.Session, error)
	GetSession(ctx context.Context, token string) (postgres.Session, error)
	UpdateSession(ctx context.Context, params *postgres.UpdateSessionParams) (postgres.Session, error)
	DeleteSession(ctx context.Context, token string) error
	GetSessionData(ctx context.Context, sess *postgres.Session, key string) (string, error)
	PutSessionData(ctx context.Context, sess *postgres.Session, key string, val string) error
	DeleteSessionData(ctx context.Context, sess *postgres.Session, key string) error
	GarbageCollect() error
}
