package web

import (
	"html/template"
	"net/http"

	"github.com/dkr290/go-projects/gonews"
	"github.com/go-chi/chi/v5"
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

	h.Get("/threads", h)

	return h
}

const threadsListHTML = `
<h1>Threads</h1>
{{range .Threads}}
    <dt><strong>{{.Title}}</strong></dt>
    <dd>{{.Description}}</dd>
{{end}}
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
