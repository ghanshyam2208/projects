package repositories

import (
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"

	"github.com/jmoiron/sqlx"
)

type RepositoryDB struct {
	sqlxClient *sqlx.DB
	appConfigs *configs.Config
}

func (r *RepositoryDB) connectDB() error {
	// Connect to the database
	sqlDb, err := sqlx.Connect("postgres", r.appConfigs.PostgresConnStr)
	if err != nil {
		return err
	}

	r.sqlxClient = sqlDb
	logger.Info("Successfully connected to the database")
	return nil
}
