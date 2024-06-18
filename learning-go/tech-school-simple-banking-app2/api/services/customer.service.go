package services

import "banking_app2/api/repositories"

type CustomerService interface {
	GetAllCustomers() ([]repositories.Customer, error)
}

type DefaultCustomerService struct {
	Repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{Repo: repo}
}

func (cs DefaultCustomerService) GetAllCustomers() ([]repositories.Customer, error) {
	return cs.Repo.FindAll()
}
