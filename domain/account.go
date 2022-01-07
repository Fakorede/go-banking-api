package domain

import (
	"go-banking-api/dto"
	"go-banking-api/errs"
)

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
