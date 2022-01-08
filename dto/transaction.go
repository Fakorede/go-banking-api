package dto

import "github.com/Fakorede/go-banking-api/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
