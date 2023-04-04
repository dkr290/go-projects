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

	m.App.Session.Put(r.Context(), "userid", "someuserid")

	render.RenderTemplate(w, "home.page.html", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	//userid := m.App.Session.GetString(r.Context(), "userid")

	// var strMap = map[string]string{
	// 	"title":  "About",
	// 	"intro":  "This page is where we talk about ourselves.",
	// 	"userid": userid,
	// }

	render.RenderTemplate(w, "about.page.html", &models.PageData{})

}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "login.page.html", &models.PageData{})

}

func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "make-post.page.html", &models.PageData{})

}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "page.page.html", &models.PageData{})

}
