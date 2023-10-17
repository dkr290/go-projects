package render

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/dkr290/go-projects/holidayapplication/pkg/config"
	"github.com/dkr290/go-projects/holidayapplication/pkg/models"
	"github.com/flosch/pongo2"
)

//////////////////////////////////////////
// func RenderTemplateNooCache(w http.ResponseWriter, tpml string, data any) {

// 	//parsedTemplate, _ := template.ParseFiles("./templates/" + tpml)
// 	parsedTemplate, err := pongo2.FromFile("./templates/" + tpml)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = parsedTemplate.ExecuteWriter(pongo2.Context{"Data": data}, w)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	// err := parsedTemplate.Execute(w, nill)
// 	// if err != nil {
// 	// 	fmt.Println("Error partsing the template ", err)
// 	// }

// }

//////////////////////////////

// render tempplate  with a cache but not ideal like  if we have a a lot of tempplates
// this can be used but ther e is better cache version

// /////////////////////////////////
// var tc = make(map[string]*pongo2.Template)

// func RenderTemplate(w http.ResponseWriter, t string, data any) {

// 	var tmpl *pongo2.Template
// 	var err error

// 	// check if we already have in the cache
// 	_, inMap := tc[t]

// 	if !inMap {
// 		// needs to create the template
// 		err := createTemplateCache(t)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		log.Println("create template and adding to the the cache")

// 	} else {
// 		// the template is in the cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.ExecuteWriter(pongo2.Context{"Data": data}, w)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// }

// func createTemplateCache(t string) error {

// 	//parse the template
// 	parsedTemplate, err := pongo2.FromFile("./templates/" + t)
// 	if err != nil {
// 		return err
// 	}

// 	// add tmp to the cache map
// 	tc[t] = parsedTemplate
// 	return nil
// }

///////////////////////////////////////////////////////////

var app *config.AppConfig

// newtemplate will set the configuration
func NewTemplate(a *config.AppConfig) {
	app = a
}

// better cache version
func RenderTemplate(w http.ResponseWriter, tpml string, data *models.TemplateData) {
	var tc map[string]*pongo2.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TempleteCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//this is to check if it is in the map
	t, ok := tc[tpml]
	if !ok {
		http.Error(w, "template is not in the cache for some resaon", http.StatusInternalServerError)
		log.Fatalln("template not in cache for some reason", ok)

	}

	// we will store the reesult of t and docble check if it is valid template
	// we create pongo context with some data
	context := pongo2.Context{
		"Title": "Sample title",
	}
	//then we execute the template and if there is no error we have a valid template
	// we are checking that whatever is stored it is executable as template because it might be not valid
	_, err := t.Execute(context)
	if err != nil {

		log.Println("error execute pongo template", err) // if the template is not valid show why it is not valid
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// default data in the templates
	data = AddDefaultData(data)

	//render the template

	err = t.ExecuteWriter(pongo2.Context{"Data": data}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func CreateTemplateCache() (map[string]*pongo2.Template, error) {

	cache := map[string]*pongo2.Template{}
	pages, err := filepath.Glob("./templates/*-page.html")
	if err != nil {
		return cache, err
	}

	// range through a slice of the pages

	for _, page := range pages {
		name := filepath.Base(page)                         // this is the template itself like just the html file name without path
		parsedTemplate, err := pongo2.FromFile("./" + page) // in this case we have parsed template for the full path
		if err != nil {
			return cache, err
		}

		cache[name] = parsedTemplate // the map cache[home-page.html] = parsedTemplate type

	}

	return cache, nil

}

// something that  will be added to all pages in this case

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}
