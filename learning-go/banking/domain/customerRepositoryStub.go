package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (customerRepositoryStub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return customerRepositoryStub.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ted", "Mumbai", "421503", "Mumbai", "true"},
		{"1002", "Robbin", "Mumbai", "421503", "Mumbai", "true"},
	}

	return CustomerRepositoryStub{customers: customers}
}
