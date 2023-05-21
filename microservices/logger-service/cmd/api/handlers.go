package main

import (
	"log"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `data:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read the json into the var

	var requestPayload JSONPayload

	_ = app.readJSON(w, r, &requestPayload)

	//insert the data

	// event := data.LogEntry{
	// 	Name: requestPayload.Name,
	// 	Data: requestPayload.Data,
	// }
	log.Println("this is from logger")

	err := app.Models.LogEntry.InsertMongo(requestPayload.Name, requestPayload.Data)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}
