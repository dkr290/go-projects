package service

import (
	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *customeerrors.AppError)
	GetCustomer(id string) (*domain.Customer, *customeerrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s *DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *customeerrors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)

}

func (s *DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *customeerrors.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepo) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repository,
	}
}
