package main

import (
	"fmt"
	"net/http"

	"github.com/dkr290/go-projects/holidayapplication/pkg/handlers"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/about", handlers.HandleAbout)

	fmt.Printf("Starting the application on port %s\n", portNumber)
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}

}
