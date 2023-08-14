package service

import "github.com/dkr290/go-projects/banking-api/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s *DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepo) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repository,
	}
}
