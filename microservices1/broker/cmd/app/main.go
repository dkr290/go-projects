package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "80"
)

type Config struct{}

func main() {

	app := Config{}
	log.Printf("Starting broker service at port %s", webPort)

	// going to the chi router in routes and then to each handler
	srv := &http.Server{

		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}