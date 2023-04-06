package main

import (
	"net/http"

	"github.com/dkr290/go-projects/web3/pkg/config"
	"github.com/dkr290/go-projects/web3/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)
	mux.Use(NoSurf)
	mux.Use(SetupSession)
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Get("/page", handlers.Repo.PageHandler)
	mux.Get("/make-post", handlers.Repo.MakePostHandler)
	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)

	return mux
}
