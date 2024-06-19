package services

import (
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/utils/errs"
)

type ICustomerService interface {
	GetAllAccounts() ([]repositories.Account, *errs.AppError)
}

type CustomerService struct {
	repo repositories.IAccountRepository
}

func (cs CustomerService) GetAllAccounts() ([]repositories.Account, *errs.AppError) {
	return cs.repo.GetAllAccounts()
}

func NewCustomerService(repository repositories.IAccountRepository) *CustomerService {
	return &CustomerService{
		repo: repository,
	}
}
