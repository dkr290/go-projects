package main

import (
	"net/http"

	"github.com/dkr290/go-projects/web3/pkg/config"
	"github.com/dkr290/go-projects/web3/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	return mux
}
