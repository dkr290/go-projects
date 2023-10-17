package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/render"
)

//////////////////////////////////////
//repository pattern
// type Repository struct {
// 	App *config.AppConfig
// }

// var Repo *Repository

// // this creates new repository
// func NewRepo(a *config.AppConfig) *Repository {
// 	return &Repository{
// 		App: a,
// 	}
// }

// func NewHandlers(r *Repository) {

// 	Repo = r
// }

////////////////////////////////////////////

/// Doing it with interfaces

type Handlers interface {
	HandleAbout(w http.ResponseWriter, r *http.Request)
	HandleHome(w http.ResponseWriter, r *http.Request)
}

type Repository struct {
	App *config.AppConfig
}

func NewHandlers(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}

}

// this is the about functions
func (m *Repository) HandleAbout(w http.ResponseWriter, r *http.Request) {
	m.App.Data.Title = "About Page"
	m.App.Description = "This is the About page"

	render.RenderTemplate(w, "about-page.html", m.App)

}

// this is the about page
func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {
	m.App.Title = "Home Page"
	m.App.Description = "This is the Home Page"

	render.RenderTemplate(w, "home-page.html", m.App)
}
