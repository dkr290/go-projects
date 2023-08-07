package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/chatws/internal/handlers"
)

func main() {

	r := routes()
	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting the webserver on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {

		log.Fatal(err)
	}
}
