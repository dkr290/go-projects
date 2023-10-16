package render

import (
	"log"
	"net/http"

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
var tc = make(map[string]*pongo2.Template)

func RenderTemplate(w http.ResponseWriter, t string, data any) {

	var tmpl *pongo2.Template
	var err error

	// check if we already have in the cache
	_, inMap := tc[t]

	if !inMap {
		// needs to create the template
		err := createTemplateCache(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("create template and adding to the the cache")

	} else {
		// the template is in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.ExecuteWriter(pongo2.Context{"Data": data}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func createTemplateCache(t string) error {

	//parse the template
	parsedTemplate, err := pongo2.FromFile("./templates/" + t)
	if err != nil {
		return err
	}

	// add tmp to the cache map
	tc[t] = parsedTemplate
	return nil
}

///////////////////////////////////////////////////////////
