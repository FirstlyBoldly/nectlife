package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func newHumaApi(mux *chi.Mux, cookieName string) huma.API {
	config := huma.DefaultConfig("NectLife API", "1.0.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"SessionCookie": {
			Type:        "apiKey",
			In:          "cookie",
			Name:        cookieName,
			Description: "Session cookie",
		},
	}
	config.Security = []map[string][]string{{"SessionCookie": {}}}

	return humachi.New(mux, config)
}
