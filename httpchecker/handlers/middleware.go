package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func WithAuthenticatedUser(c *fiber.Ctx) error {

	log.Println("this is getting called")
	return c.Next()
}
