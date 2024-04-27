package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type ICustomerRepo interface {
	FindAll() ([]Customer, error)
}
