package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/wyllow-xyz/wyllow/internal/components/pages"
)

// Creates and returns a new HTTP server with a predefined router.
func New() *http.Server {
	router := chi.NewMux()

	router.Get("/", handleHome)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}

// Renders the home page
func handleHome(w http.ResponseWriter, r *http.Request) {
	component := pages.Home()
	component.Render(r.Context(), w)
}
