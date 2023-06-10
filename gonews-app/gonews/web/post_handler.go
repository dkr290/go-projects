package web

import (
	"html/template"
	"net/http"

	"github.com/dkr290/go-projects/gonews"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type PostHandler struct {
	store gonews.Store
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread gonews.Thread
		CSRF   template.HTML
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
	tmpl.Execute(w, data{Thread: t, CSRF: csrf.TemplateField(r)})
}

func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread   gonews.Thread
		Post     gonews.Post
		Comments []gonews.Comment
		CSRF     template.HTML
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

	tmpl.Execute(w, data{Thread: t, Post: p, Comments: cc, CSRF: csrf.TemplateField(r)})
}

func (h *PostHandler) Store(w http.ResponseWriter, r *http.Request) {

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
func (h *PostHandler) Vote(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "postID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.store.Post(id)
	if err != nil {
		http.Error(w, "error getting single post from the database: "+err.Error(), http.StatusNotFound)
		return
	}
	dir := r.URL.Query().Get("dir")

	if dir == "up" {
		p.Votes++

	} else if dir == "down" {
		p.Votes--
	}

	if err := h.store.UpdatePost(&p); err != nil {

		http.Error(w, "error updating comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusFound)

}
