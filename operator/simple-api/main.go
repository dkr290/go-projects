package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown"
	}
	if _, err := fmt.Fprintf(w, "<h1 style='font-size:24px;'>Version: %s</h1>\n", version); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := ":9090"
	fmt.Println("Server running on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
