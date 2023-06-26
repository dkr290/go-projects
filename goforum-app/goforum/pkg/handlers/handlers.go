package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/forms"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/render"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/repository"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/repository/dbrepo"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// repository used for the different handlers
var Repo *Repository

func NewRepo(ac *config.AppConfig, db *pgx.Conn) *Repository {

	return &Repository{
		App: ac,
		DB:  dbrepo.NewPostgresRepo(db, ac),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {

	var artList models.ArticleList

	artList, err := m.DB.GetThreeArticles()
	if err != nil {
		log.Println(err)
		return
	}

	for i := range artList.Content {
		fmt.Println(artList.Content[i])
	}

	m.App.Session.Put(r.Context(), "user_id", "someuser")
	data := make(map[string]any)
	data["artList"] = artList

	log.Println(data)

	render.RenderTemplate(w, r, "home.html", &models.PageData{
		Data: data,
	})

}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "about.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	log.Println("test")

	render.RenderTemplate(w, r, "login.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) MakePost(w http.ResponseWriter, r *http.Request) {

	// if user is not logged in redirect to the login
	if !m.App.Session.Exists(r.Context(), "user_id") {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	var emptyArticle models.Article
	data := make(map[string]any)
	data["article"] = emptyArticle

	render.RenderTemplate(w, r, "make-post.html", &models.PageData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostMakePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uID := (m.App.Session.Get(r.Context(), "user_id")).(int)

	article := models.Post{
		Title:   r.Form.Get("blog_title"),
		Content: r.Form.Get("blog_article"),
		UserID:  uID,
	}

	form := forms.New(r.PostForm)

	if article.Title == "" {
		form.FormNoValueError("blog_title")

	} else {
		form.MinLenght("blog_title", "4")
	}

	article.Content = strings.TrimSpace(article.Content)
	if len(article.Content) == 0 {
		form.FormNoValueError("blog_article")

	} else {
		form.MinLenght("blog_article", "8")
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

	// write to the database
	log.Println("about ot insert to db")
	err = Repo.DB.InsertPost(article)
	if err != nil {
		log.Fatal(err)
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
	data := make(map[string]any)
	data["article"] = article

	render.RenderTemplate(w, r, "article-received.html", &models.PageData{
		Data: data,
	})

}

func (m *Repository) Page(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)

	render.RenderTemplate(w, r, "page.html", &models.PageData{
		StrMap: strMap,
	})
}

func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {

	_ = m.App.Session.RenewToken(r.Context())
	log.Println("from post Login")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	if len(email) == 0 {
		form.FormNoValueError("email")
	}
	if len(password) == 0 {
		form.FormNoValueError("password")
	}
	form.ValidateEmail(email)

	if !form.Valid() {

		data := make(map[string]any)

		data["email"] = email
		data["password"] = password
		render.RenderTemplate(w, r, "login.html", &models.PageData{
			Form: form,
			Data: data,
		})

		return
	}

	id, _, err := m.DB.AuthenticateUser(email, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "invalid email or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "Valid Login")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (m *Repository) LogOutHandler(w http.ResponseWriter, r *http.Request) {

	_ = m.App.Session.Destroy(r.Context())

	_ = m.App.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
