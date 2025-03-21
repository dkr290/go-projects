package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "supersecret1" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}
	time.Sleep(10 * time.Second)
	msg := "Hello client"
	w.Write([]byte(msg))
}

func main() {
	fmt.Println("starting the server")
	log.Fatal(http.ListenAndServe(":8080", &server{}))
}
