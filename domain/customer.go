package domain

import (
	"go-banking-api/dto"
	"go-banking-api/errs"
)

type CustomerRepository interface {
	// status - 0, 1, ""
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(id string) (*Customer, *errs.AppError)
}

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
