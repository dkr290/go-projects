package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/gonews/postgres"
	"github.com/dkr290/go-projects/gonews/web"
)

func main() {

	store, err := postgres.NewStore("postgres://postgres:password@postgres/news?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)

	http.ListenAndServe(":3000", h)

}
