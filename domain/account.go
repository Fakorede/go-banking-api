package domain

import (
	"github.com/Fakorede/go-banking-api/dto"
	"github.com/Fakorede/go-banking-api/errs"
)

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

type Account struct {
	AccountId   string	`db:"account_id"`
	CustomerId  string	`db:"customer_id"`
	OpeningDate string	`db:"opening_date"`
	AccountType string	`db:"account_type"`
	Amount      float64	`db:"amount"`
	Status      string	`db:"status"`
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
