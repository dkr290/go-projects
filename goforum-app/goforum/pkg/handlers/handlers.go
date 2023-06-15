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

	render.RenderTemplate(w, r, "home.html", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "about.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "login.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) MakePost(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "make-post.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) PostMakePost(w http.ResponseWriter, r *http.Request) {

	blog_title := r.Form.Get("blog_title")
	blog_article := r.Form.Get("blog_article")

	w.Write([]byte(blog_title))
	w.Write([]byte(blog_article))
}

func (m *Repository) Page(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "page.html", &models.PageData{
		StrMap: strMap,
	})
}
