package services

import (
	"go-banking-api/domain"
	"go-banking-api/dto"
	"go-banking-api/errs"

	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewDefaultAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDTO()

	return &response, nil
}
