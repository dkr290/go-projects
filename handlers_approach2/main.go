package main

import (
	"database/sql"
	"log"
	"net/http"
	"prod/db"
	"prod/handlers"
	"prod/helpers"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewMux()

	var dd *sql.DB
	d := db.Mysql{
		DB: dd,
	}

	router.Get("/", helpers.MakeHandlers(handlers.HandleIndex, &d))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
