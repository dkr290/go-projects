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
		r.Post("/{id}/delete", h.ThreadsDelete)

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

const threadsListHTML = `
<h1>Threads</h1>
<dl>
{{range .Threads}}
    <dt><strong>{{.Title}}</strong></dt>
    <dd>{{.Description}}</dd>
	<dd>
     <form action="/threads/{{.ID}}/delete" method="POST">
	 <button type="submit">Delete</button>
	 </form>
	</dd>
{{end}}

<a href="/threads/new">Create thread</a>
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
