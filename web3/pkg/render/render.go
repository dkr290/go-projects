package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/web3/models"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, pd *models.PageData) {

	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Template in cache")
	}

	tmpl = tmplCache[t]
	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
	}
}

func makeTemplateCache(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.html",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tmplCache[t] = tmpl
	return nil

}