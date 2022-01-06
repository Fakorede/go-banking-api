package domain

import "go-banking-api/errs"

type Customer struct {
	ID          string	`db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// status - 0, 1, ""
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(id string) (Customer, *errs.AppError)
}
