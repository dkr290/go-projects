package main

import (
	"log"
	"net/http"
)

func main() {

	r := routes()
	log.Println("Starting the webserver on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {

		log.Fatal(err)
	}
}
