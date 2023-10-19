package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/models"
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

	RemoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	pageData := make(map[string]string)
	pageData["Title"] = "About page"
	pageData["Description"] = "This is about page"
	pageData["remote_ip"] = RemoteIP

	render.RenderTemplate(w, "about-page.html", &models.TemplateData{
		StringMap: pageData,
	})

}

// this is the about page
func (m *Repository) HandleHome(w http.ResponseWriter, r *http.Request) {

	RemoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", RemoteIP)

	pageData := make(map[string]string)
	pageData["Title"] = "Home page"
	pageData["Description"] = "This is home page"

	render.RenderTemplate(w, "home-page.html", &models.TemplateData{
		StringMap: pageData,
	})
}
