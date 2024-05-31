package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "ghanshyam", "thane", "421503", "1993-08-22", "active"},
		{"1002", "kiran", "thane", "421503", "1993-08-22", "active"},
	}
	return CustomerRepositoryStub{customers}
}
