package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nectgrams-webapp-team/nectlife/internal/config"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type SessionService struct {
	db *postgres.Queries
}

func NewSessionStore(db *postgres.Queries) SessionServiceInterface {
	return &SessionService{db: db}
}

func NewSessionToken() (string, error) {
	id := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(id), nil
}

func GenerateExpiresAtTime() time.Time {
	return time.Now().UTC().Add(time.Duration(config.Data.EXPIRES_AT_IN_HOURS) * time.Hour)
}

func (s *SessionService) CreateSession(ctx context.Context) (postgres.Session, error) {
	token, err := NewSessionToken()
	if err != nil {
		return postgres.Session{}, err
	}

	return s.db.CreateSession(ctx, postgres.CreateSessionParams{
		Token: token,
		Data:  pgtype.Hstore{},
		ExpiresAt: pgtype.Timestamptz{
			Time:  GenerateExpiresAtTime(),
			Valid: true,
		},
	})
}

func (s *SessionService) GetSession(ctx context.Context, token string) (postgres.Session, error) {
	return s.db.GetSession(ctx, token)
}

func (s *SessionService) UpdateSession(ctx context.Context, params *postgres.UpdateSessionParams) (postgres.Session, error) {
	return s.db.UpdateSession(ctx, *params)
}

func (s *SessionService) DeleteSession(ctx context.Context, token string) error {
	return s.db.DeleteSession(ctx, token)
}

func (s *SessionService) GetSessionData(ctx context.Context, sess *postgres.Session, key string) (string, error) {
	val, ok := sess.Data[key]
	if !ok {
		text := fmt.Sprintf("failed to get session data with key: %s", key)
		return "", errors.New(text)
	}

	return *val, nil
}

func (s *SessionService) PutSessionData(ctx context.Context, sess *postgres.Session, key string, val string) error {
	sess.Data[key] = &val
	_, err := s.db.UpdateSession(
		ctx,
		postgres.UpdateSessionParams{
			ID:        sess.ID,
			Token:     sess.Token,
			Data:      sess.Data,
			ExpiresAt: sess.ExpiresAt,
		},
	)
	return err
}

func (s *SessionService) DeleteSessionData(ctx context.Context, sess *postgres.Session, key string) error {
	delete(sess.Data, key)
	if _, err := s.db.UpdateSession(
		ctx,
		postgres.UpdateSessionParams{
			ID:        sess.ID,
			Token:     sess.Token,
			Data:      sess.Data,
			ExpiresAt: sess.ExpiresAt,
		},
	); err != nil {
		return err
	}

	return nil
}

func (s *SessionService) GarbageCollect() error {
	return s.db.GarbageCollect(context.Background())
}
