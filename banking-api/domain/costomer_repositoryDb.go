package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepoDb struct {
	client *sql.DB
}

func (c *CustomerRepoDb) FindAll() ([]Customer, error) {

	findAllSQL := `SELECT customer_id, name, date_of_birth, city, zipcode, status
				   FROM customers;`
	rows, err := c.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while quering customer table", err)
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.DateOfBirth,
			&c.City,
			&c.Zipcode,
			&c.Status,
		); err != nil {
			log.Println("Error when scanning the customers", err)
			return nil, err
		}

		customers = append(customers, c)

	}

	return customers, nil

}

func (c *CustomerRepoDb) ById(id string) (*Customer, error) {
	SQL := `SELECT customer_id, name, date_of_birth, city, zipcode, status
			FROM customers where customer_id = ?;`

	row := c.client.QueryRow(SQL, id)
	var cus Customer

	if err := row.Scan(
		&cus.Id,
		&cus.Name,
		&cus.DateOfBirth,
		&cus.City,
		&cus.Zipcode,
		&cus.Status,
	); err != nil {
		log.Println("Error when scanning the customer", err)
		return nil, err
	}

	return &cus, nil
}

func NewCustomerRepoDb() *CustomerRepoDb {
	client, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	err = testDb(client)
	if err != nil {
		log.Fatal(err)
	}

	return &CustomerRepoDb{
		client: client,
	}
}

func testDb(client *sql.DB) error {
	counts := 0

	for {
		err := client.Ping()
		if err != nil {
			log.Println("Mysql server is not yet ready")
			counts++
		} else {
			log.Println("*** Pinged database successfully! ***")
			return nil
		}
		if counts > 10 {
			log.Println("Error connection to the database", err)
			return err
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
