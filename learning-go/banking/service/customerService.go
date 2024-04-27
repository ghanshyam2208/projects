package service

import "github.com/gnbaviskar/bankingApp/domain"

type ICustomerService interface {
	findAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	customerRepo domain.ICustomerRepo
}

func (defaultCustomerService DefaultCustomerService) findAllCustomers() ([]domain.Customer, error) {
	return defaultCustomerService.customerRepo.FindAll()
}

func NewCustomerService(customerRepo domain.ICustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{customerRepo: customerRepo}
}
