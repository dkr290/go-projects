package domain

import (
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/dto"
)

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

func (c Customer) ToDto() dto.CustomerResponse {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	resp := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      statusAsText,
	}

	return resp
}
