package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func errorCheck(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func writeHelper(writer http.ResponseWriter, msg string) {
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

func getSum(x, y int) int {
	return x + y
}

func devideFunc(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	}
	return x / y, nil
}

func devideHandler(w http.ResponseWriter, r *http.Request) {
	var x, y float64 = 28.8, 2.49
	div, err := devideFunc(x, y)
	if err != nil {
		log.Fatal(err)
	}

	writeHelper(w, fmt.Sprintf("The division of %f / %f = %f", x, y, div))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var x, y int = 4, 4
	writeHelper(w, "Hello internet")
	sum := getSum(x, y)
	output := fmt.Sprintf("\t%d + %d = %d\n", x, y, sum)
	writeHelper(w, output)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getsum", addHandler)
	http.HandleFunc("/division", devideHandler)

	log.Println("Server started")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)

	}

}
