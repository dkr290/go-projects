package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-fuego/fuego"
)

func (app *application) routes() *fuego.Server {
	s := fuego.NewServer(fuego.WithGlobalMiddlewares(middleware.Recoverer))

	// registr

	// static access

	return s
}
