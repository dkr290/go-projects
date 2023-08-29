package main

import (
	"github.com/dkr290/go-projects/banking-api/cmd/app"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
