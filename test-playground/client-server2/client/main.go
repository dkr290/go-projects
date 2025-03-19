package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturn(msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)

	r, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	message := messageData{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}

func main() {
	msg := messageData{
		Message: "Hi server",
	}
	data := postDataAndReturn(msg)
	fmt.Println(data)
}
