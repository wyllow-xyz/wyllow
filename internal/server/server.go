package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New() *http.Server {
	router := chi.NewMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}
