package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnREsponse(filename string) string {
	fileDataBuffer := bytes.Buffer{}
	multiPartWriter := multipart.NewWriter(&fileDataBuffer)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	formFile, err := multiPartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}

	multiPartWriter.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080", &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := postFileAndReturnREsponse("test.txt")
	fmt.Println(data)
}
