package domain

import "github.com/ghanshyam2208/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
