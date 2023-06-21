package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	//connect to the db

	count := 0
	var conn *pgx.Conn
	var err error
	connString := "postgres://postgres:password@localhost:5432/blog_db"

	for {

		//conn, err = sql.Open("pgx", "host=postgres port=5432 dbname=blog_db user=postgres password=password")
		conn, err = pgx.Connect(context.Background(), connString)
		if err != nil && count <= 10 {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			count++
			time.Sleep(10 * time.Second)
		} else if err == nil {
			log.Println("Connected to the database")
			break
		} else if count > 10 {
			log.Fatalln("Too many retries to connect to the database")
			break
		}

	}

	defer conn.Close(context.Background())
}
