package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Rsvp struct {
	Name, Email, Phone string
	WillATtend         bool
}
type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)
var templ = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTemplates() {

	templNames := [5]string{"list", "form", "sorry", "welcome", "thanks"}
	for i, name := range templNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err != nil {
			panic(err)
		} else {
			templ[name] = t
			fmt.Println("Loaded template", i, name)
		}
	}
}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templ["welcome"].Execute(w, nil)
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	templ["list"].Execute(w, nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ["form"].Execute(w, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})
	}
}
