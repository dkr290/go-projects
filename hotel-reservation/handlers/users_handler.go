package handlers

import (
	"github.com/dkr290/go-projects/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {

	u := types.User{
		FirstName: "James",
		LastName:  "Ath the watercooler",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
