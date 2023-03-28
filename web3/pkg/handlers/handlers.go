package handlers

import (
	"net/http"

	"github.com/dkr290/go-projects/web3/models"
	"github.com/dkr290/go-projects/web3/pkg/config"
	"github.com/dkr290/go-projects/web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

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
	render.RenderTemplate(w, "home.page.html", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	var strMap = map[string]string{
		"title": "About",
		"intro": "This page is where we talk about ourselves.",
	}

	render.RenderTemplate(w, "about.page.html", &models.PageData{
		StrMap: strMap,
	})

}
