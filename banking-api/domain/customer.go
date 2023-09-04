package domain

import "github.com/dkr290/go-projects/banking-api/pkg/customeerrors"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

// repository for the customers
type CustomerRepo interface {
	// status == 1 active status == 0 inactive status == ""
	FindAll(status string) ([]Customer, *customeerrors.AppError)
	ById(id string) (*Customer, *customeerrors.AppError)
}
