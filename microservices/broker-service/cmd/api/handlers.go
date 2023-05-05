package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//we need to define predictable json that is supposed to be passed

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {

	var rp RequestPayload
	//read the json that we are receiving and checking for error
	if err := app.readJSON(w, r, &rp); err != nil {
		app.errorJSON(w, err)
		return
	}

	//actions based on the json we have received

	switch rp.Action {
	case "auth":
		app.authenticate(w, rp.Auth)

	default:
		app.errorJSON(w, errors.New("unknown action"))

	}

}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json well send to the auth microservice

	jsonData, _ := json.MarshalIndent(a, "", "\t")

	//call the service

	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))

	if err != nil {

		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	//make sure we get back the correct statuscode

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create variable will read response  body into

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {

		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)

}
