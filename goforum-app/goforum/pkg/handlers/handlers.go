package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/forms"
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

	render.RenderTemplate(w, r, "make-post.html", &models.PageData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostMakePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	article := models.Article{
		BlogTitle:   r.Form.Get("blog_title"),
		BlogArticle: r.Form.Get("blog_article"),
	}

	form := forms.New(r.PostForm)

	if article.BlogTitle == "" {
		form.FormNoValueError("blog_title")
	}

	article.BlogArticle = strings.TrimSpace(article.BlogArticle)
	if len(article.BlogArticle) == 0 {
		form.FormNoValueError("blog_article")
	}
	if !form.Valid() {
		data := make(map[string]any)
		data["article"] = article
		render.RenderTemplate(w, r, "make-post.html", &models.PageData{
			Form: form,
			Data: data,
		})
		return
	}

}

func (m *Repository) Page(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "page.html", &models.PageData{
		StrMap: strMap,
	})
}
