package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type jsonResponse struct {
	Message string
}

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {

	log.Println("calling broker")
	var data any
	if err := app.Helpers.ReadJsonFromHttp(w, r, &data); err != nil {
		log.Fatal(errors.New("unable to read json from http request"))
		return
	}

	jsonData, _ := json.MarshalIndent(data, "", "\t")
	registerServiceURl := "http://register"
	request, err := http.NewRequest("POST", registerServiceURl, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		panic(err)

	}

	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {

		log.Fatalln(err)
		return
	}

	app.Helpers.SendJsonResponse(w, http.StatusAccepted, jsonFromService)

}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("Authenticate controller")
}
