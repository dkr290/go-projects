package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	type requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var rp requestPayload

	if err := app.readJSON(w, r, &rp); err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validate user against the database

	user, err := app.Models.User.GetByEmail(rp.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(rp.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
