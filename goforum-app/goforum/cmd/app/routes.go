package main

import (
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//handle routing for the application
//we will do chi routing

func routes(app *config.AppConfig) http.Handler {

	// Mux = r from router

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(LogRequestInfo)
	r.Use(NoSurf)
	r.Use(SetupSession)
	r.Get("/", handlers.Repo.HomeHandler)
	r.Get("/about", handlers.Repo.AboutHandler)
	r.Get("/login", handlers.Repo.Login)
	r.Get("/makepost", handlers.Repo.MakePost)
	r.Post("/makepost", handlers.Repo.PostMakePost)
	r.Post("/login", handlers.Repo.PostLogin)

	r.Get("/article-received", handlers.Repo.ArticleReceived)

	r.Get("/login", handlers.Repo.Page)

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r

}
