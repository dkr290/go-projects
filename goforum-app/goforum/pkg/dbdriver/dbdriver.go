package dbdriver

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

// hold the database connection
// allow us to connect to multiply databases
var conn *pgx.Conn

const (
	maxOpenDbConnections = 20
	maxIdleDbConnections = 10
	maxDBLifeTime        = 5 * time.Minute
)

func ConnectDatabase(dsn string) *pgx.Conn {
	count := 0
	var err error

	for {

		conn, err = pgx.Connect(context.Background(), dsn)
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
	conn.Config().ConnectTimeout = maxDBLifeTime

	defer conn.Close(context.Background())

	return conn
}
