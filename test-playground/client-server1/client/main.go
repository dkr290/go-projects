package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Names struct {
	Names []string `json:"names"`
}

func main() {
	err := getDataAndParseResponce()
	if err != nil {
		log.Println(err)
	}
}

func getDataAndParseResponce() error {
	message := Names{}
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		return fmt.Errorf("err %s", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	err = json.Unmarshal(data, &message)
	if err != nil {
		return fmt.Errorf("error %v", err)
	}
	electricCount := 0
	boogalloCount := 0
	for _, v := range message.Names {
		if v == "Electric" {
			electricCount += 1
		}
		if v == "Boogaloo" {
			boogalloCount += 1
		}
	}
	fmt.Println("Electric Count", electricCount)
	fmt.Println("Boogallo Count", boogalloCount)
	return nil
}
