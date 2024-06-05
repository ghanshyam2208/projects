package domain

import (
	"github.com/ghanshyam2208/banking/dto"
	"github.com/ghanshyam2208/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) getStatusText() string {
	statusText := "inactive"
	if c.Status == "1" {
		statusText = "active"
	}
	return statusText
}

func (c Customer) CreateCustomerResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.getStatusText(),
	}
}
