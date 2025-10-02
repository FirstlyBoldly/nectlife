package session

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/nectgrams-webapp-team/nectlife/internal/postgres"
)

type sessionResponseWriter struct {
	http.ResponseWriter
	sessionHandler *SessionHandler
	request        *http.Request
	done           bool
}

func (w *sessionResponseWriter) Write(b []byte) (int, error) {
	writeCookieIfNecessary(w)
	return w.ResponseWriter.Write(b)
}

func (w *sessionResponseWriter) WriteHeader(code int) {
	writeCookieIfNecessary(w)
	w.ResponseWriter.WriteHeader(code)
}

func (w *sessionResponseWriter) UnWrap() http.ResponseWriter {
	return w.ResponseWriter
}

func writeCookieIfNecessary(w *sessionResponseWriter) {
	if w.done {
		return
	}

	sess, ok := w.request.Context().Value(sessCtxKey).(*postgres.Session)
	if !ok {
		slog.Error("failed to get session in the request context")
		return
	}

	maxAge := int(time.Until(sess.ExpiresAt.Time) / time.Second)
	cookie := &http.Cookie{
		Name:     w.sessionHandler.cookieName,
		Value:    sess.Token,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  sess.ExpiresAt.Time,
		MaxAge:   maxAge,
	}

	http.SetCookie(w.ResponseWriter, cookie)
	w.done = true
}
