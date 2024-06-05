package service

import (
	"github.com/ghanshyam2208/banking/domain"
	"github.com/ghanshyam2208/banking/dto"
	"github.com/ghanshyam2208/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.CreateCustomerResponse()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
