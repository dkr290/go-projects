package main

import (
	"log"

	"github.com/dkr290/go-projects/banking-api/cmd/app"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")

	}

	logger.Info("Starting the application...")
	app.Start()
}
