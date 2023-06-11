package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/go-projects/gonews"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
)

type ThreadHandler struct {
	store    gonews.Store
	sessions *scs.SessionManager
}

func (h *ThreadHandler) List() http.HandlerFunc {
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

func (h *ThreadHandler) Create(w http.ResponseWriter, r *http.Request) {

	type data struct {
		CSRF template.HTML
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/thread_create.html"))
	tmpl.Execute(w, data{CSRF: csrf.TemplateField(r)})
}

func (h *ThreadHandler) Show(w http.ResponseWriter, r *http.Request) {

	type data struct {
		Thread gonews.Thread
		Posts  []gonews.Post
		CSRF   template.HTML
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
	tmpl.Execute(w, data{Thread: thred, Posts: pp, CSRF: csrf.TemplateField(r)})
}

func (h *ThreadHandler) Store(w http.ResponseWriter, r *http.Request) {

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

func (h *ThreadHandler) Delete(w http.ResponseWriter, r *http.Request) {

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
