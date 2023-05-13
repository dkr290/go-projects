package main

import (
	"log"
	"net/http"
)

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("connecter to register")
}
