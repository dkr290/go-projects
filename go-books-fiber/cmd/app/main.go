package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dkr290/go-projects/go-books-fiber/models"
	"github.com/dkr290/go-projects/go-books-fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

const webPort = "3000"

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Config struct {
	DB *gorm.DB
}

func (c *Config) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", c.CreateBook)
	api.Delete("delete_book/:id", c.DeleteBook)
	api.Get("/get_books/:id", c.GetBookById)
	api.Get("/books", c.GetBooks)

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.PgConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("cloud not lod the database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}

	c := Config{
		DB: db,
	}

	app := fiber.New()
	c.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", webPort)))

}
