package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/handlers"
)

const webPort = "8080"

func main() {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil); err != nil {
		log.Fatal(err)
	}

}
