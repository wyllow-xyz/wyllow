package server

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/wyllow-xyz/wyllow/internal/components/pages"
)

//go:embed assets/static/*
var staticFiles embed.FS

// Creates and returns a new HTTP server with a predefined router.
func New() *http.Server {
	router := chi.NewMux()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	// Serve embedded static files
	staticFileServer := http.FileServer(http.FS(staticFiles))
	router.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/assets/static" + r.URL.Path
		staticFileServer.ServeHTTP(w, r)
	}))

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
