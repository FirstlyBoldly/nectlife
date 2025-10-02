package nlerrs

import (
	"log/slog"
	"net/http"
)

// Logs the error `text` and sends a 500 error to the client
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request, text string) {
	slog.Error(text)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error: " + text + "\n"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	text := "URL \"" + r.URL.Path + "\" does not exist"
	slog.Error(text)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found: " + text + "\n"))
}
