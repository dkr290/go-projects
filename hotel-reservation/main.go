package main

import (
	"flag"

	"github.com/dkr290/go-projects/hotel-reservation/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "Listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", handlers.HandleGetUsers)
	apiv1.Get("/user/:id", handlers.HandleGetUser)
	app.Listen(*listenAddr)

}
