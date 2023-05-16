package main

import (
	"errors"
	"fmt"
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

type jsonResponse struct {
	Message string
}

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

	payload := jsonResponse{
		Message: fmt.Sprintf("User with email %s and id %d is created", ap.Email, userID),
	}

	app.Helpers.SendJsonResponse(w, http.StatusAccepted, payload)

}
