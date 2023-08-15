package domain

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
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, error)
}
