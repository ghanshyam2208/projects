package repositories

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type DefaultCustomerRepository struct {
	Customer []Customer
}

func (s DefaultCustomerRepository) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepository() CustomerRepository {
	customers := []Customer{
		{"1001", "ghanshyam", "thane", "421503", "1993-08-22", "active"},
		{"1002", "kiran", "thane", "421503", "1993-08-22", "active"},
	}
	return DefaultCustomerRepository{Customer: customers}
}
