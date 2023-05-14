package main

import (
	"fmt"
	"log"
	"net/http"
	"register/data"
)

const webPort = "80"

type Config struct {
	Models data.Models
}

func main() {

	app := Config{}
	log.Printf("Starting register service at port %s", webPort)

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
