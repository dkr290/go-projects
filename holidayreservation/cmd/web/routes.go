package main

import (
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/handlers"
	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	h := handlers.NewHandlers(app)
	render.NewTemplate(app)

	//trying with pat
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(h.HandleHome))
	// mux.Get("/about", http.HandlerFunc(h.HandleAbout))

	// using chi router

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(hitLogger)
	mux.Use(noSurf)
	mux.Use(sessionLoad)

	mux.Get("/", http.HandlerFunc(h.HandleHome))
	mux.Get("/about", http.HandlerFunc(h.HandleAbout))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
