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
	h.Route("/threads", func(r chi.Router) {

		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate)
		r.Post("/", h.ThreadsStore)

	})

	return h
}

const threadsListHTML = `
<h1>Threads</h1>
<dl>
{{range .Threads}}
    <dt><strong>{{.Title}}</strong></dt>
    <dd>{{.Description}}</dd>
{{end}}
</dl>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []gonews.Thread
	}
	tmpl := template.Must(template.New("").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		thread, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: thread})
	}
}

const threadCreateHTML = `
<h1>New Thread</h1>
<form action="/threads" method="POST">
   <table>
       <tr>
                <td>Title</td>
				<td><input type="text" name="title" /></td>

	   </tr>

	   <tr>
                <td>Description</td>
				<td><input type="text" name="description" /></td>
				
	   </tr>

   </table>
   <button type="submit"> Create  Thread</button>

</form>
`

func (h *Handler) ThreadsCreate(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("").Parse(threadCreateHTML))
	tmpl.Execute(w, nil)
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
