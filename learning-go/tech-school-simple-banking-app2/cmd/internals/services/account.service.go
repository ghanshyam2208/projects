package services

import (
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/utils/errs"
)

type IAccountService interface {
	GetAllAccounts() ([]repositories.Account, *errs.AppError)
}

type AccountService struct {
	repo repositories.IAccountRepository
}

func (cs AccountService) GetAllAccounts() ([]repositories.Account, *errs.AppError) {
	return cs.repo.GetAllAccounts()
}

func NewAccountService(repository repositories.IAccountRepository) *AccountService {
	return &AccountService{
		repo: repository,
	}
}
