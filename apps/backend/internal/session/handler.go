package session

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type ctxKey int

var sessCtxKey ctxKey

type SessionHandler struct {
	store      SessionServiceInterface
	cookieName string
}

func (h *SessionHandler) Start(r *http.Request) (*postgres.Session, *http.Request) {
	var session postgres.Session
	cookie, err := r.Cookie(h.cookieName)
	if err == nil {
		session, err = h.store.GetSession(r.Context(), cookie.Value)
		if err != nil {
			slog.Error("failed to get session: " + err.Error())
		}
	}

	if err != nil || !h.ValidateSession(r.Context(), &session) {
		session, err = h.store.CreateSession(r.Context())
		if err != nil {
			slog.Error("failed to create a session: " + err.Error())
			return nil, r
		}
	}

	sessionPtr := &session
	ctx := context.WithValue(r.Context(), sessCtxKey, sessionPtr)
	r = r.WithContext(ctx)

	return sessionPtr, r
}

func (h *SessionHandler) ValidateSession(ctx context.Context, session *postgres.Session) bool {
	if time.Now().UTC().Compare(session.ExpiresAt.Time) >= 0 {
		if err := h.store.DeleteSession(ctx, session.Token); err != nil {
			slog.Error("failed to delete invalid session: " + err.Error())
		}

		return false
	}

	return true
}

func (h *SessionHandler) Migrate(ctx context.Context, session *postgres.Session) error {
	token, err := NewSessionToken()
	if err != nil {
		slog.Error("failed to create new session token: " + err.Error())
		return err
	}

	sess, err := h.store.UpdateSession(ctx, &postgres.UpdateSessionParams{
		ID:    session.ID,
		Token: token,
		Data:  session.Data,
		ExpiresAt: pgtype.Timestamptz{
			Time:  GenerateExpiresAtTime(),
			Valid: true,
		},
	})
	if err != nil {
		slog.Error("failed to update session token: " + err.Error())
		return err
	}

	*session = sess
	return nil
}

func (h *SessionHandler) GarbageCollect(d time.Duration) {
	ticker := time.NewTicker(d)
	for range ticker.C {
		err := h.store.GarbageCollect()
		if err != nil {
			slog.Error("failed to garbage collect: " + err.Error())
		}
	}
}

func (h *SessionHandler) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, rws := h.Start(r)
		sessionWriter := &sessionResponseWriter{
			ResponseWriter: w,
			sessionHandler: h,
			request:        rws,
		}

		w.Header().Add("Vary", "Cookie")
		w.Header().Add("Cache-Control", `no-cache="Set-Cookie"`)

		next.ServeHTTP(sessionWriter, rws)

		writeCookieIfNecessary(sessionWriter)
	})
}

func GetRequestSession(ctx context.Context) (*postgres.Session, bool) {
	session, ok := ctx.Value(sessCtxKey).(*postgres.Session)
	if !ok {
		slog.Error("failed to get session from the request context")
		return nil, ok
	}

	return session, ok
}

func NewSessionHandler(
	interval time.Duration,
	store SessionServiceInterface,
	cookieName string,
) SessionHandlerInterface {
	manager := SessionHandler{
		store:      store,
		cookieName: cookieName,
	}

	go manager.GarbageCollect(interval)

	return &manager
}
