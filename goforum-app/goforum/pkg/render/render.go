package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
)

// making the template cache
var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, pageData *models.PageData) {

	var tmpl *template.Template
	var err error

	_, inMap := templateCache[t]
	if !inMap {
		makeTemplateCache(t)
	} else {

		fmt.Println("Template is in cache already")
	}

	tmpl = templateCache[t]
	err = tmpl.Execute(w, pageData)

	if err != nil {
		log.Println(err)
	}

}

func makeTemplateCache(t string) {

	//one entry for each template to render
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/layout.html",
	}

	tmpl := template.Must(template.ParseFiles(templates...))
	templateCache[t] = tmpl

}
