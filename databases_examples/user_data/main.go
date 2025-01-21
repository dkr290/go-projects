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
		{Id: 1, Name: "Szabo Daniel", Email: "daniel@packt.com"},
		{Id: 2, Name: "Szabo Florian", Email: "florian@packt.com"},
	}
	conf := db.InitConfig()
	d, err := db.InitDb(conf, 10)
	if err != nil {
		log.Fatal(err)
	}

	database := db.PsqlDatabase{
		Db: d,
	}
	err = database.CreateTables()
	if err != nil {
		log.Fatal(err)
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
		fmt.Printf("Name: %s with email: %s\n", v.Name, v.Email)
	}
	defer database.Db.Close()
}
