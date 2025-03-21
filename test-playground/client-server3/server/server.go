package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadFile, UploadFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadFile.Close()
	fileContent, err := io.ReadAll(uploadFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("./%s", UploadFileHeader.Filename), fileContent, 0600)
	if err != nil {
		log.Fatal(err)
	}
	var buff []byte
	w.Write(fmt.Appendf(buff, "%s uploaded", UploadFileHeader.Filename))
}

func main() {
	fmt.Println("Starting the server")
	log.Fatal(http.ListenAndServe(":8080", &server{}))
}
