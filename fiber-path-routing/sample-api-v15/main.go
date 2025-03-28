package main

import (
	"fiber-path-routing/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	// Create a new Fiber app with templates
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// Add logger middleware
	app.Use(logger.New())
	// Serve static files (for CSS, JS, etc.)
	app.Static("/", "./public")
	h := handlers.New()
	// sample group

	// Home API Endpoints
	app.Get("/", h.IndexHandler)
	app.Get("/home", h.HomeHandler)
	app.Get("/health", h.HealthHandler)
	app.Get("/status", h.StatusHandler)
	app.Get("/metrics", h.MetricsHandler)
	app.Get("/docs", h.DocsHandler)
	app.Get("/redoc", h.RedocHandler)

	// Predict API group under p
	predict := app.Group("/v15")
	predict.Get("/", h.PredictIndexHandler)
	predict.Get("/home", h.PredictHomeHandler)
	predict.Get("/docs", h.PredictDocsHandler)
	predict.Get("/redoc", h.PredictRedocHandler)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
