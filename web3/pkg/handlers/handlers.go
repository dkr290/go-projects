package handlers

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/web3/models"
	"github.com/dkr290/go-projects/web3/pkg/config"
	"github.com/dkr290/go-projects/web3/pkg/forms"
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

	render.RenderTemplate(w, r, "home.page.html", &models.PageData{})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	//userid := m.App.Session.GetString(r.Context(), "userid")

	// var strMap = map[string]string{
	// 	"title":  "About",
	// 	"intro":  "This page is where we talk about ourselves.",
	// 	"userid": userid,
	// }

	render.RenderTemplate(w, r, "about.page.html", &models.PageData{})

}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "login.page.html", &models.PageData{})

}

func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {

	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	render.RenderTemplate(w, r, "make-post.page.html", &models.PageData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// blog_title := r.Form.Get("blog_title")
	// blog_article := r.Form.Get("blog_article")

	// fmt.Fprintf(w, "%s", blog_title)
	// fmt.Fprintf(w, "%s", blog_article)

	article := models.Article{
		BlogTitle:   r.Form.Get("blog_title"),
		BlogArticle: r.Form.Get("blog_article"),
	}

	form := forms.New(r.PostForm)

	form.HasRequired("blog_title", "blog_article")

	form.MinLenght("blog_title", 5, r)
	form.MinLenght("blog_article", 5, r)

	data := make(map[string]interface{})
	if !form.Valid() {

		data["article"] = article
		render.RenderTemplate(w, r, "make-post.page.html", &models.PageData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)

}

func (m *Repository) ArticleReceived(w http.ResponseWriter, r *http.Request) {
	article, ok := m.App.Session.Get(r.Context(), "article").(models.Article)
	if !ok {
		log.Println("Can't get data from session")

		m.App.Session.Put(r.Context(), "error", "Can't get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	data := make(map[string]interface{})
	data["article"] = article

	render.RenderTemplate(w, r, "article-received.page.html", &models.PageData{

		Data: data,
	})

}

func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "page.page.html", &models.PageData{})

}
