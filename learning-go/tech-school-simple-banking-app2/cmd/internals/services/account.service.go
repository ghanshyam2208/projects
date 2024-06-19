package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/utils/errs"
)

type IAccountService interface {
	GetAllAccounts() ([]dto.AccountDto, *errs.AppError)
}

type AccountService struct {
	repo repositories.IAccountRepository
}

func (cs AccountService) GetAllAccounts() ([]dto.AccountDto, *errs.AppError) {
	repoAccounts, err := cs.repo.GetAllAccounts()
	if err != nil {
		return nil, err
	}
	response := make([]dto.AccountDto, 0)
	for _, account := range repoAccounts {
		response = append(response, account.CreateAccountResponse())
	}

	return response, nil
}

func NewAccountService(repository repositories.IAccountRepository) *AccountService {
	return &AccountService{
		repo: repository,
	}
}
