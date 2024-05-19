package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Rsvp struct {
	Name, Email, Phone string
	Willattend         bool
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
	templ["list"].Execute(w, responses)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ["form"].Execute(w, formData{
			Rsvp: &Rsvp{
				Willattend: false,
			},
			Errors: []string{},
		})
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		respData := Rsvp{
			Name:       r.Form["name"][0],
			Email:      r.Form["email"][0],
			Phone:      r.Form["phone"][0],
			Willattend: r.Form["willattend"][0] == "true",
		}
		errors := []string{}
		if respData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if respData.Email == "" {
			errors = append(errors, "Please enter your email address")
		}
		if respData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templ["form"].Execute(w, formData{
				Rsvp: &respData, Errors: errors,
			})
		} else {

			responses = append(responses, &respData)
			if respData.Willattend {
				templ["thanks"].Execute(w, respData.Name)
			} else {
				templ["sorry"].Execute(w, respData.Name)
			}
		}
	}
}
