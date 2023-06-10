package web

import (
	"html/template"
	"net/http"

	"github.com/dkr290/go-projects/gonews"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type Handler struct {
	*chi.Mux
	store gonews.Store
}

func NewHandler(store gonews.Store, csrfKey []byte) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	var threads = ThreadHandler{store: store}
	var posts = PostHandler{store: store}

	h.Use(middleware.Logger)
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false)))
	h.Get("/", h.Home)
	h.Route("/threads", func(r chi.Router) {

		r.Get("/", threads.List())
		r.Get("/new", threads.Create)
		r.Post("/", threads.Store)
		r.Get("/{id}", threads.Show)
		r.Post("/{id}/delete", threads.Delete)
		r.Get("/{id}/new", posts.Create)
		r.Post("/{id}", posts.Store)
		r.Get("/{threadID}/{postID}", posts.Show)
		r.Get("/{threadID}/{postID}/vote", posts.Vote)
		r.Post("/{threadID}/{postID}", h.CommentsStore)

	})
	h.Get("/comments/{id}/vote", h.CommentsVote)

	// h.Route("/html", func(r chi.Router) {
	// 	r.Get("/", h.HtmlGet)
	// })

	return h
}

// func (h *Handler) HtmlGet(w http.ResponseWriter, r *http.Request) {

// 	templ := template.Must(template.New("layout.html").ParseGlob("templates-playground/includes/*.html"))
// 	templ = template.Must(templ.ParseFiles("templates-playground/layout.html", "templates-playground/childtemplate.html"))

// 	type prarams struct {
// 		Title   string
// 		Text    string
// 		Lines   []string
// 		Number1 int
// 		Number2 int
// 	}

// 	templ.Execute(w, prarams{
// 		Title: "News site",
// 		Text:  "Welcome to our go news site",
// 		Lines: []string{
// 			"Line1",
// 			"Line2",
// 			"Line3",
// 		},
// 		Number1: 12,
// 		Number2: 123,
// 	})

// }

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Posts []gonews.Post
	}

	pp, err := h.store.Posts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	tmpl.Execute(w, data{Posts: pp})
}

func (h *Handler) CommentsStore(w http.ResponseWriter, r *http.Request) {

	content := r.FormValue("content")

	idStr := chi.URLParam(r, "postID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := &gonews.Comment{
		ID:      uuid.New(),
		PostID:  id,
		Content: content,
	}

	if err := h.store.CreateComment(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)

}

func (h *Handler) CommentsVote(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := h.store.Comment(id)
	if err != nil {
		http.Error(w, "error getting single comment from the database: "+err.Error(), http.StatusNotFound)
		return
	}
	dir := r.URL.Query().Get("dir")

	if dir == "up" {
		c.Votes++

	} else if dir == "down" {
		c.Votes--
	}

	if err := h.store.UpdateComment(&c); err != nil {

		http.Error(w, "error updating comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)

}
