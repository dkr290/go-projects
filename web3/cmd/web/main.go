package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/web3/pkg/config"
	"github.com/dkr290/go-projects/web3/pkg/handlers"
)

func main() {

	var app config.AppConfig
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: routes(&app),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
