package domain

import "go-banking-api/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(id string) (Customer, *errs.AppError)
}
