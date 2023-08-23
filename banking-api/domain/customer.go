package domain

import "github.com/dkr290/go-projects/banking-api/pkg/customeerrors"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// repository for the customers
type CustomerRepo interface {
	FindAll() ([]Customer, *customeerrors.AppError)
	ById(id string) (*Customer, *customeerrors.AppError)
}
