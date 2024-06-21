package repositories

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/utils/configs"
	"fmt"
	"time"
)

type TransferRepositoryDb struct {
	appConfigs *configs.Config
}

func (t *TransferRepositoryDb) TransferAmount(payload dto.TransferAmountDto) (Transfer, error) {
	fmt.Println("wiring")
	return Transfer{
		Id:            1,
		FromAccountId: payload.FromAccountId,
		ToAccountId:   payload.ToAccountId,
		Amount:        payload.Amount,
		CreatedAt:     time.Now(),
	}, nil
}

func NewTransferRepo(configs *configs.Config) *TransferRepositoryDb {
	repo := &TransferRepositoryDb{
		appConfigs: configs,
	}
	// repo.connectDB()
	return repo
}
