package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	Item string
}

func IndexHandler(db *sql.DB, c *fiber.Ctx) error {

	var res string
	var todos []string

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
		log.Fatalln(err)
		c.JSON("An error selectiong todos occured")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	defer rows.Close()

	return c.Render("index", fiber.Map{"Todos": todos})
}

func PostHandler(db *sql.DB, c *fiber.Ctx) error {

	newTodo := todo{}

	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured %v", err)
		return c.SendString(err.Error())
	}

	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
	return c.Redirect("/")

}

func DeleteHandler(db *sql.DB, c *fiber.Ctx) error {

	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}

func UpdateHandler(db *sql.DB, c *fiber.Ctx) error {

	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newitem, olditem)
	return c.Redirect("/")
}
