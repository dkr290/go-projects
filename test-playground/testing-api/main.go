package main

import (
	"embed"
	"html/template"
	"log"
	"testing-api/internal/cmiddleware"
	"testing-api/internal/config"
	"testing-api/internal/handlers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-fuego/fuego"
)

//go:embed templates
var templateFS embed.FS

func main() {
	newIpMiddleware := cmiddleware.New()

	app := config.New(newIpMiddleware)
	routes(app)
}

func routes(app *config.Application) {
	tmpls := template.Must(template.ParseFS(templateFS,
		"templates/*.page.html",
	))

	s := fuego.NewServer(
		fuego.WithGlobalMiddlewares(middleware.Recoverer, app.CMiddlewares.AddIpToContext),
		fuego.WithTemplateFS(templateFS),
		fuego.WithTemplates(tmpls),
	)
	h := handlers.New(app)
	fuego.Get(s, "/", h.Home)

	log.Println("Starting the server on port :9999")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
