package service

import (
	"time"

	"github.com/dkr290/go-projects/banking-api/domain"
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/dto"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *customeerrors.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepo
}

func (s *DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *customeerrors.AppError) {

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil

}
