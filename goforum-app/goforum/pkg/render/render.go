package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
	"github.com/justinas/nosurf"
)

// making the template cache
var templateCache = make(map[string]*template.Template)

var app *config.AppConfig

func NewAppConfig(a *config.AppConfig) {
	app = a
}

func AddCSRFData(pd *models.PageData, r *http.Request) *models.PageData {

	pd.CSRFToken = nosurf.Token(r)

	if app.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd

}

func RenderTemplate(w http.ResponseWriter, r *http.Request, t string, pageData *models.PageData) {

	var tmpl *template.Template
	var err error

	_, inMap := templateCache[t]
	if !inMap {
		makeTemplateCache(t)
	} else {

		fmt.Println("Template is in cache already")
	}

	tmpl = templateCache[t]
	pd := AddCSRFData(pageData, r)
	err = tmpl.Execute(w, pd)

	if err != nil {
		log.Println("error template execution", err)
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
