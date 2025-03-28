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
	app.Static("/static", "./public")
	h := handlers.New()
	// sample group
	sample := app.Group("/sample/v15")

	// Home API Endpoints
	sample.Get("/", h.IndexHandler)
	sample.Get("/home", h.HomeHandler)
	sample.Get("/health", h.HealthHandler)
	sample.Get("/status", h.StatusHandler)
	sample.Get("/metrics", h.MetricsHandler)
	sample.Get("/docs", h.DocsHandler)
	sample.Get("/redoc", h.RedocHandler)

	// Predict API group under sample
	predict := sample.Group("/v15")
	predict.Get("/", h.PredictIndexHandler)
	predict.Get("/home", h.PredictHomeHandler)
	predict.Get("/docs", h.PredictDocsHandler)
	predict.Get("/redoc", h.PredictRedocHandler)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
