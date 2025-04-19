package main

import "log"

type application struct{}

func main() {
	app := application{}
	s := app.routes()

	log.Println("Starting the server on port :9999")
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
