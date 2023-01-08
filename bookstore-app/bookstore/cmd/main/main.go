package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/bookstore-app/bookstore/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))

}
