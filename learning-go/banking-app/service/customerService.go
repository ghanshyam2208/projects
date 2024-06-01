package service

import "github.com/ghanshyam2208/banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	return s.repo.ById((id))
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
