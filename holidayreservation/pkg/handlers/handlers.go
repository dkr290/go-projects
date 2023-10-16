package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// this creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {

	Repo = r
}

type Data struct {
	Title string
}

// this is the about functions
func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "About Page",
	}

	render.RenderTemplate(w, "about-page.html", data)

}

// this is the about page
func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "Home Page",
	}
	render.RenderTemplate(w, "home-page.html", data)
}
