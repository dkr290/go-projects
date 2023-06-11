package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/gonews/postgres"
	"github.com/dkr290/go-projects/gonews/web"
)

func main() {

	dsn := "postgres://postgres:password@postgres/news?sslmode=disable"
	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	sessions, err := web.NewSessionsManager(dsn)
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("eQaAy7NR/Ju9mR1+trtn1ojV9S7AKmKIlxknL/LpRY2ugqFAx6C69GlV8hgdy+9p")

	h := web.NewHandler(store, sessions, csrfKey)

	http.ListenAndServe(":3000", h)

}
