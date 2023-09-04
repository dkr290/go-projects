package service

import (
	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/dto"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *customeerrors.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *customeerrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (s *DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *customeerrors.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	cc, err := s.repo.FindAll(status)
	var cusResp []dto.CustomerResponse
	if err != nil {
		return nil, err
	}
	for _, c := range cc {
		cusResp = append(cusResp, c.ToDto())
	}

	return cusResp, nil

}

func (s *DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *customeerrors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepo) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repository,
	}
}
