package main

import (
	"net/http"

	"github.com/dkr290/go-projects/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
)

func (c *Config) CreateBook(context *fiber.Ctx) error {

	book := Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err

	}

	err = c.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been added"})

	return nil
}

func (c *Config) DeleteBook(context *fiber.Ctx) error {
	books := []models.Books{}
	err := c.DB.Find(&books).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "coud not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books find successfully",
		"data":    books,
	})

	return nil

}

func (c *Config) GetBooks(context *fiber.Ctx) error {

}
