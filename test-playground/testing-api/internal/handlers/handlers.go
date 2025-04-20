package handlers

import (
	"log"
	"testing-api/internal/config"

	"github.com/go-fuego/fuego"
)

type Handlers struct {
	app *config.Application
}

func New(app *config.Application) *Handlers {
	return &Handlers{
		app: app,
	}
}

func (h *Handlers) Home(c fuego.ContextNoBody) (fuego.CtxRenderer, error) {
	dataIP := h.app.CMiddlewares.GetIpFromContext(c.Context())
	render, err := c.Render("home.page.html", fuego.H{
		"IP": dataIP,
	})
	if err != nil {
		// Handle the error appropriately, e.g., log it and return an internal server error
		log.Printf("Error rendering template: %v", err)
		return nil, err // Or perhaps c.Status(http.StatusInternalServerError).String("Internal Server Error")
	}
	return render, nil
}
