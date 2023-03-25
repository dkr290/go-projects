package main

import (
	"log"
	"net/http"
	"text/template"
)

func errorCheck(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err, "error from template parsing")
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err, "execute template")

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")

}

func main() {

	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
