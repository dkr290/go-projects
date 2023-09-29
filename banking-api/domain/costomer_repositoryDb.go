package domain

import (
	"database/sql"

	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepoDb struct {
	client *sqlx.DB
}

func (c *CustomerRepoDb) FindAll(status string) ([]Customer, *customeerrors.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := `SELECT customer_id, name, date_of_birth, city, zipcode, status
				   FROM customers;`

		err = c.client.Select(&customers, findAllSQL)
		//rows, err = c.client.Query(findAllSQL) ## because of sqlx

	} else {
		findAllSQL := `SELECT customer_id, name, date_of_birth, city, zipcode, status
				   FROM customers where status = ?;`

		err = c.client.Select(&customers, findAllSQL, status)
		//rows, err = c.client.Query(findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error quering customers table " + err.Error())
		return nil, customeerrors.NewUnexpectedError("Error while quering customer table")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, customeerrors.NewUnexpectedError("Unexpected database error")
	// }
	// for rows.Next() {
	// 	var c Customer
	// 	if err := rows.Scan(
	// 		&c.Id,
	// 		&c.Name,
	// 		&c.DateOfBirth,
	// 		&c.City,
	// 		&c.Zipcode,
	// 		&c.Status,
	// 	); err != nil {
	// 		return nil, customeerrors.NewUnexpectedError("Error scanning the customers")
	// 	}

	//	customers = append(customers, c)

	return customers, nil

}

func (c *CustomerRepoDb) ById(id string) (*Customer, *customeerrors.AppError) {
	SQL := `SELECT customer_id, name, date_of_birth, city, zipcode, status
			FROM customers where customer_id = ?;`

	//row := c.client.QueryRow(SQL, id) old way of doing it without sqlx
	var cus Customer

	err := c.client.Get(&cus, SQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customeerrors.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error when scanning the customer" + err.Error())
			return nil, customeerrors.NewUnexpectedError("unexpected database error")
		}
	}

	// if err := row.Scan(
	// 	&cus.Id,
	// 	&cus.Name,
	// 	&cus.DateOfBirth,
	// 	&cus.City,
	// 	&cus.Zipcode,
	// 	&cus.Status,
	// ); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, customeerrors.NewNotFoundError("customer not found")
	// 	}
	// 	logger.Error("Error when scanning the customer" + err.Error())
	// 	return nil, customeerrors.NewUnexpectedError("unexpected database error")

	// }

	return &cus, nil
}

func NewCustomerRepoDb(dbClient *sqlx.DB) *CustomerRepoDb {

	return &CustomerRepoDb{
		client: dbClient,
	}
}
