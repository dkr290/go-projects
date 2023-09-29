package domain

import (
	"strconv"

	"github.com/dkr290/go-projects/banking-api/pkg/customeerrors"
	"github.com/dkr290/go-projects/banking-api/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepoDb struct {
	client *sqlx.DB
}

func (a *AccountRepoDb) Save(account Account) (*Account, *customeerrors.AppError) {

	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) values(?,?,?,?,?)"

	result, err := a.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, customeerrors.NewUnexpectedError("unexpected error from the database")
	}

	accId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inerted id for the new account: " + err.Error())
		return nil, customeerrors.NewUnexpectedError("unexpected error from the database")
	}
	account.AccountId = strconv.FormatInt(accId, 10)

	return &account, nil

}

func NewAccountRepoDb(dbClient *sqlx.DB) *AccountRepoDb {

	return &AccountRepoDb{
		client: dbClient,
	}

}
