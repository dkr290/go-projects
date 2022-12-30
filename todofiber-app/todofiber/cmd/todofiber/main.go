package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/template/html"

	"github.com/dkr290/go-devops/todofiber/handlers"
	"github.com/dkr290/go-devops/todofiber/repository"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var Repo = repository.NewRepo()

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	Repo.DbUser = os.Getenv("DATABASE_USER")
	Repo.DbHost = os.Getenv("DATABASE_HOST")
	Repo.DbName = "todos"
	Repo.DbPass = os.Getenv("DATABASE_PASS")
	Repo.APPPort = os.Getenv("APP_PORT")
	Repo.DbPort = os.Getenv("DATABASE_PORT")

	if Repo.DbPort == "" {
		Repo.DbPort = "5432"
	}

	connInfo := fmt.Sprintf("user=%s password=%s port=%s host=%s sslmode=disable", Repo.DbUser, Repo.DbPass, Repo.DbPort, Repo.DbHost)
	initdb, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	dbName := Repo.DbName
	_, err = initdb.Exec("create database " + dbName)

	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			//handle the error
			fmt.Println("Database already created")

		} else {
			log.Fatalln("Cannot create database", err)
		}
	}

	initdb.Close()

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=disable", Repo.DbUser, Repo.DbPass, Repo.DbHost, Repo.DbPort, dbName)
	fmt.Println(connStr)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//Then execute your query for creating table
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS todos( %s )", "item text"))

	if err != nil {

		log.Fatal("Create table failed ", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(db, c)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return handlers.PostHandler(db, c)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return handlers.UpdateHandler(db, c)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return handlers.DeleteHandler(db, c)
	})

	if Repo.APPPort == "" {
		Repo.APPPort = "3000"
	}

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", Repo.APPPort)))
}
