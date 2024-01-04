package main

import (
	"github.com/LobovVit/url-shortener/internal/app/config"
	"github.com/LobovVit/url-shortener/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg := config.GetConfig()
	mux := chi.NewRouter()
	mux.Post("/", handlers.SetShortHandler)
	mux.Get("/{key}", handlers.GetShortHandler)

	return http.ListenAndServe(cfg.Host, mux)
}
