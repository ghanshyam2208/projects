package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/utils/errs"
)

type IAccountService interface {
	GetAllAccounts(int, int) ([]dto.AccountDto, *errs.AppError)
	CreateAccount(dto.CreateAccountDto) (dto.AccountDto, *errs.AppError)
}

type AccountService struct {
	repo repositories.IAccountRepository
}

func (cs AccountService) GetAllAccounts(page int, pageSize int) ([]dto.AccountDto, *errs.AppError) {
	repoAccounts, err := cs.repo.GetAllAccounts(page, pageSize)
	if err != nil {
		return nil, err
	}
	response := make([]dto.AccountDto, 0)
	for _, account := range repoAccounts {
		response = append(response, account.CreateAccountResponse())
	}

	return response, nil
}

func (cs AccountService) CreateAccount(createAccountDto dto.CreateAccountDto) (dto.AccountDto, *errs.AppError) {
	account, err := cs.repo.CreateAccount(createAccountDto)
	if err != nil {
		return dto.AccountDto{}, err
	}
	return account.CreateAccountResponse(), nil
}

func NewAccountService(repository repositories.IAccountRepository) *AccountService {
	return &AccountService{
		repo: repository,
	}
}
