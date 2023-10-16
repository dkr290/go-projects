package render

import (
	"net/http"
	"path/filepath"

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

// better cache version
func RenderTemplate(w http.ResponseWriter, tpml string, data any) {

	//create template cache

	tc, err := createTemplateCache()

	err = parsedTemplate.ExecuteWriter(pongo2.Context{"Data": data}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func createTemplateCache() (map[string]*pongo2.Template, error) {

	cache := map[string]*pongo2.Template{}
	pages, err := filepath.Glob("./templates/*-page.html")
	if err != nil {
		return cache, err
	}

	// range through a slice of the pages

	for _, page := range pages {
		name := filepath.Base(page)                  // this is the template itself like just the html file name without path
		parsedTemplate, err := pongo2.FromFile(page) // in this case we have parsed template for the full path
		if err != nil {
			return cache, err
		}

		cache[name] = parsedTemplate // the map cache[home-page.html] = parsedTemplate type

	}

	return cache, nil

}
