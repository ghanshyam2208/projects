package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
)

type ITransferService interface {
	TransferAmount(dto.CreateAccountDto) (*repositories.Transfer, error)
}

type DefaultTransferService struct {
	repo repositories.ITransferRepo
}

func (d DefaultTransferService) TransferAmount(payload dto.TransferAmountDto) (*repositories.Transfer, error) {
	return d.repo.TransferAmount(payload)
}

func NewDefaultTransferService(repo repositories.ITransferRepo) *DefaultTransferService {
	return &DefaultTransferService{
		repo: repo,
	}
}
