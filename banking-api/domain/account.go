package domain

import (
	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/dto"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepo interface {
	Save(account Account) (*Account, *customeerrors.AppError)
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
