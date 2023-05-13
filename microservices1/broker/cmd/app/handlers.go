package main

import "net/http"

type AuthPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (app *Config) Register(w http.ResponseWriter, a AuthPayload) {

}
