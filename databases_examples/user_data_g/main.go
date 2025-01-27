package main

import (
	"fmt"
	"log"
	"userdata/internal/db"
	"userdata/internal/models"

	_ "github.com/lib/pq"
)

func main() {
	users := []models.User{
		{FirstName: "Jao", LastName: "Daniel", Email: "daniel@packt.com"},
		{FirstName: "Szlao", LastName: "Florian", Email: "florian@packt.com"},
	}
	conf := db.InitConfig()
	d, err := db.InitDb(conf, 10)
	if err != nil {
		log.Fatal(err)
	}

	database := db.PsqlDatabase{
		Db: d,
	}
	err = database.InsertUsers(users)
	if err != nil {
		log.Fatal(err)
	}

	u, err := database.GetAllRecords()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range u {
		fmt.Printf("FirstName: %s, Lastname: %s with email: %s\n", v.FirstName, v.LastName, v.Email)
	}
}
