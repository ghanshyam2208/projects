package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
)

type IAccountService interface {
	GetAllAccounts(int, int) ([]dto.AccountDto, error)
	CreateAccount(dto.CreateAccountDto) (dto.AccountDto, error)
	UpdateAccount(dto.UpdateAccountDto) error
	DeleteAccount(int64) error
}

type AccountService struct {
	repo repositories.IAccountRepository
}

func (cs AccountService) GetAllAccounts(page int, pageSize int) ([]dto.AccountDto, error) {
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

func (cs AccountService) CreateAccount(createAccountDto dto.CreateAccountDto) (dto.AccountDto, error) {
	account, err := cs.repo.CreateAccount(createAccountDto)
	if err != nil {
		return dto.AccountDto{}, err
	}
	return account.CreateAccountResponse(), nil
}

func (cs AccountService) UpdateAccount(updateAccountDto dto.UpdateAccountDto) error {
	return cs.repo.UpdateAccount(updateAccountDto.Id, updateAccountDto)
}

func (cs AccountService) DeleteAccount(id int64) error {
	return cs.repo.DeleteAccount(id)
}

func NewAccountService(repository repositories.IAccountRepository) *AccountService {
	return &AccountService{
		repo: repository,
	}
}
