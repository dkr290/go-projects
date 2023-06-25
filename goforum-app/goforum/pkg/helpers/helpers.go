package helpers

import (
	"log"
	"net/http"

	"github.com/dkr290/go-projects/goforum-app/goforum/pkg/config"
)

var app config.AppConfig

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsAuthenticated(r *http.Request) bool {

	exists := app.Session.Exists(r.Context(), "user_id")

	return exists
}
