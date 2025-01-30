package main

import (
	"log"
	"net/http"
	"server-templ/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MakeHandler(handlers.IndexHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
