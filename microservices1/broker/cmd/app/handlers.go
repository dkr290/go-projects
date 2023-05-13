package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AuthPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (app *Config) Register(w http.ResponseWriter, r *http.Request) {

	var ap AuthPayload
	jsonData, _ := json.MarshalIndent(ap, "", "\t")
	registerServiceURl := "http://register"
	request, err := http.NewRequest("POST", registerServiceURl, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {

		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		panic(err)

	}

}
