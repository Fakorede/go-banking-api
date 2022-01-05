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
	// status - 0, 1, ""
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(id string) (Customer, *errs.AppError)
}
