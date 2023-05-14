package main

import (
	"encoding/json"
	"errors"
	"io"
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

	maxBytes := 1048576 //one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&ap); err != nil {
		panic(err)

	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		log.Fatal(errors.New("body must have only a single JSON value"))

	}

	//insert user

	userID, err := app.Models.User.Insert(ap.Email, ap.Password, ap.FirstName, ap.LastName)
	if err != nil {
		log.Fatal(errors.New("error crating user"))
		return
	}

	log.Printf("The user with the id %d\n", userID)

}
