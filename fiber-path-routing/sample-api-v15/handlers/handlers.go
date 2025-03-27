package handlers

import "github.com/gofiber/fiber/v2"

type Handlers struct{}

func New() *Handlers {
	return &Handlers{}
}

// Home API Handlers
func (h *Handlers) IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Sample API Home",
	})
}

func (h *Handlers) HomeHandler(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Sample Dashboard",
	})
}

func (h *Handlers) HealthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Service is healthy",
	})
}

func (h *Handlers) StatusHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "operational",
		"message": "All systems normal",
	})
}

func (h *Handlers) MetricsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"metrics": fiber.Map{
			"requests":   1024,
			"uptime":     "99.9%",
			"response":   "32ms",
			"throughput": "1.2k req/s",
		},
	})
}

func (h *Handlers) DocsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Swagger UI documentation",
		"url":     "https://swagger.io",
	})
}

func (h *Handlers) RedocHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ReDoc UI documentation",
		"url":     "https://redoc.ly",
	})
}

// Predict API Handlers
func (h *Handlers) PredictIndexHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Predict API Index",
		"version": "v15",
		"endpoints": fiber.Map{
			"home":  "/sample/v15/home",
			"docs":  "/sample/v15/docs",
			"redoc": "/sample/v15/redoc",
		},
	})
}

func (h *Handlers) PredictHomeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Predict API Home",
		"version": "v15",
	})
}

func (h *Handlers) PredictDocsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Predict API Swagger UI",
		"version": "v15",
		"url":     "https://swagger.io",
	})
}

func (h *Handlers) PredictRedocHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Predict API ReDoc UI",
		"version": "v15",
		"url":     "https://redoc.ly",
	})
}
