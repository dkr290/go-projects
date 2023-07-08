package main

import (
	"context"
	"flag"

	"log"

	"github.com/dkr290/go-projects/hotel-reservation/db"
	"github.com/dkr290/go-projects/hotel-reservation/handlers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "reservations"
	userColl = "users"
)

func main() {

	listenAddr := flag.String("listenAddr", ":5000", "Listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// handlers initializations
	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))
	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	app.Listen(*listenAddr)

}
