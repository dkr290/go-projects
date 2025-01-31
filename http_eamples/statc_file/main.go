package main

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type Customer struct {
	Name string
}

type Templ struct {
	tpl *template.Template
}

func (h *Templ) HandleIndex(w http.ResponseWriter, r *http.Request) error {
	param := r.URL.Query()
	cust := Customer{}

	name, ok := param["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}

	return h.tpl.Execute(w, cust)
}

func MakeHandler(next func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func NewTempl(tplPath string) (*Templ, error) {
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &Templ{tpl: tpl}, nil
}

func main() {
	http.Handle(
		"/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))),
	)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./templates/index.html")
	// })

	tpl, err := NewTempl("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", MakeHandler(tpl.HandleIndex))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
