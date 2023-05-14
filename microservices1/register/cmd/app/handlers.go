package main

import (
	"errors"
	"log"
	"net/http"
)

type authPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var ap authPayload

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("connected to register")

	if err := app.Helpers.ReadJsonFromHttp(w, r, &ap); err != nil {
		log.Fatal(errors.New("unable to read json from http request"))
		return
	}

	//insert user

	userID, err := app.Models.User.Insert(ap.Email, ap.Password, ap.FirstName, ap.LastName)
	if err != nil {
		log.Fatal(errors.New("error crating user"))
		return
	}

	log.Printf("The user with the id %d\n", userID)

}
