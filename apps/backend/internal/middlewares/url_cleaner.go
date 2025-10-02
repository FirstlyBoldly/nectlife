package middlewares

import (
	"net/http"
	"path/filepath"
)

func URLCleaner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = filepath.Clean(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
