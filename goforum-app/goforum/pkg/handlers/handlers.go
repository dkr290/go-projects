package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// repository used for the different handlers
var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {

	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {

	m.App.Session.Put(r.Context(), "userid", "someuser")

	render.RenderTemplate(w, "home.html", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About us"
	strMap["intro"] = "This page is where we talk about ourselves this is like that"

	userid := m.App.Session.GetString(r.Context(), "userid")
	strMap["userid"] = userid

	render.RenderTemplate(w, "about.html", &models.PageData{
		StrMap: strMap,
	})
}
