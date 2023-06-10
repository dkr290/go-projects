package web

import (
	"html/template"
	"net/http"

	"github.com/dkr290/go-projects/gonews"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Handler struct {
	*chi.Mux
	store gonews.Store
}

func NewHandler(store gonews.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Get("/", h.Home)
	h.Route("/threads", func(r chi.Router) {

		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate)
		r.Post("/", h.ThreadsStore)
		r.Get("/{id}", h.ThreadsShow)
		r.Post("/{id}/delete", h.ThreadsDelete)
		r.Get("/{id}/new", h.PostsCreate)
		r.Post("/{id}", h.PostsStore)
		r.Get("/{threadID}/{postID}", h.PostsShow)
		r.Post("/{threadID}/{postID}", h.CommentsStore)

	})

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

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
	tmpl.Execute(w, nil)
}

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []gonews.Thread
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/threads.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		thread, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: thread})
	}
}

func (h *Handler) ThreadsCreate(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread_create.html"))
	tmpl.Execute(w, nil)
}

func (h *Handler) ThreadsShow(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread gonews.Thread
		Posts  []gonews.Post
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread.html"))
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//query the threads by its id
	thred, err := h.store.Thread(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pp, err := h.store.PostsByThread(thred.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.Execute(w, data{Thread: thred, Posts: pp})
}

func (h *Handler) ThreadsStore(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	description := r.FormValue("description")

	if err := h.store.CreateThread(&gonews.Thread{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/threads", http.StatusFound)

}

func (h *Handler) ThreadsDelete(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.store.DeleteThread(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/threads", http.StatusFound)
}

func (h *Handler) PostsCreate(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread gonews.Thread
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/post_create.html"))

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := h.store.Thread(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data{Thread: t})
}

func (h *Handler) PostsShow(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread   gonews.Thread
		Post     gonews.Post
		Comments []gonews.Comment
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/post.html"))

	postIdStr := chi.URLParam(r, "postID")

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	threadIdStr := chi.URLParam(r, "threadID")

	threadId, err := uuid.Parse(threadIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := h.store.Thread(threadId)
	if err != nil {
		http.Error(w, "error Getting thread from the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	p, err := h.store.Post(postId)
	if err != nil {
		http.Error(w, "error getting post from the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	cc, err := h.store.CommentsByPost(p.ID)
	if err != nil {
		http.Error(w, "error getting comments from the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data{Thread: t, Post: p, Comments: cc})
}

func (h *Handler) PostsStore(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	content := r.FormValue("content")

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := h.store.Thread(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := &gonews.Post{
		ID:       uuid.New(),
		ThreadID: t.ID,
		Title:    title,
		Content:  content,
	}

	if err := h.store.CreatePost(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/threads/"+t.ID.String()+"/"+p.ID.String(), http.StatusFound)

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
