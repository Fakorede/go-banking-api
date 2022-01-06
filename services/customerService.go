package services

import (
	"go-banking-api/domain"
	"go-banking-api/dto"
	"go-banking-api/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	var customers []dto.CustomerResponse

	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	for _, v := range c {
		customers = append(customers, v.ToDTO())
	}

	return customers, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDTO()

	return &response, nil
}
