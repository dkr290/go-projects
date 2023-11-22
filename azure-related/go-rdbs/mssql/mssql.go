package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	server := os.Args[1]
	user := os.Args[2]
	password := os.Args[3]
	port, _ := strconv.Atoi(os.Args[4])
	database := os.Args[5]
	query := os.Args[6]

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var db *sql.DB
	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Querying table
	count, err := Select(db, query)
	if err != nil {
		log.Fatal("Error querying table: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

}

// Select reads all records
func Select(db *sql.DB, query string) (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf(query)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var column string

		// Get values from row.
		err := rows.Scan(&column)
		if err != nil {
			return -1, err
		}

		fmt.Printf("Row No. %d: %s\n", count+1, column)
		count++
	}

	return count, nil
}
