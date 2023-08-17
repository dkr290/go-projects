package service

import (
	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *customeerrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s *DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s *DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *customeerrors.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepo) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repository,
	}
}
