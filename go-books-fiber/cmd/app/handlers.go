package main

import (
	"fmt"
	"net/http"

	"github.com/dkr290/go-projects/go-books-fiber/models"
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
	books := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := c.DB.Delete(&books, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "coud not delete book"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted successfully",
	})

	return nil

}

func (c *Config) GetBooks(context *fiber.Ctx) error {
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

func (c *Config) GetBookById(context *fiber.Ctx) error {

	id := context.Params("id")
	book := &models.Books{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is ", id)
	if err := c.DB.Where("id = ?", id).First(&book).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id found sucessfully",
		"data":    book,
	})

	return nil
}
