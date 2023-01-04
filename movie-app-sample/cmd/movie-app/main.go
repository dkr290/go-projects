package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/movie-app-sample/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// new router in mux
	r := mux.NewRouter()

	r.HandleFunc("/movies", handlers.GetMovies()).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie()).Methods("GET")
	r.HandleFunc("/movies/", handlers.CreateMovie()).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie()).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie()).Methods("DELETE")

	fmt.Printf("Starting the server at port %s", "8080")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))

}
