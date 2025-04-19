package main

import (
	"embed"
	"html/template"
	"log"
	"testing-api/internal/config"
	"testing-api/internal/handlers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-fuego/fuego"
)

//go:embed templates
var templateFS embed.FS

func main() {
	app := config.New()
	s := routes(app)

	log.Println("Starting the server on port :9999")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func routes(app *config.Application) *fuego.Server {
	tmpls := template.Must(template.ParseFS(templateFS,
		"templates/*.page.html",
	))
	s := fuego.NewServer(
		fuego.WithGlobalMiddlewares(middleware.Recoverer),
		fuego.WithTemplateFS(templateFS),
		fuego.WithTemplates(tmpls),
	)
	h := handlers.New(app)

	// register routes
	fuego.Get(s, "/", h.Home)
	// static access

	return s
}
