package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"
	randomedata "banking_app2/cmd/utils/randomeData"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var accountService *AccountService
var accountRepoDB *repositories.AccountRepositoryDB

// TestMain is the entry point for testing, used for setup and teardown
func TestMain(m *testing.M) {
	// Setup before all tests
	setup()

	// Run tests
	code := m.Run()

	// Teardown after all tests
	teardown()

	// Exit with the proper code
	os.Exit(code)
}

// setup initializes resources before all tests
func setup() {
	rootDir, stdErr := filepath.Abs(filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir("__FILE__"))), "../../../"))
	if stdErr != nil {
		logger.Error("could not load root dir " + stdErr.Error())
	}
	appConfig, stdErr := configs.LoadConfig(rootDir, "app-test")
	if stdErr != nil {
		logger.Error("could not load configs " + stdErr.Error())
	}

	// Initialize the service with a repository
	accountRepoDB = repositories.NewAccountsRepo(appConfig)
	accountService = NewAccountService(accountRepoDB)
}

// teardown cleans up resources after all tests
func teardown() {
	accountRepoDB.CleanAccounts()
}

// SetupTest initializes resources before each test
func SetupTest() {
	// Clean the database or repository before each test
	accountRepoDB.CleanAccounts()
}

// TeardownTest cleans up resources after each test
func TeardownTest() {
	// Optionally clean up after each test if needed
}

func getRandomAcc() dto.CreateAccountDto {
	return dto.CreateAccountDto{
		Owner:    randomedata.RandomOwner(),
		Balance:  randomedata.RandomInt(100, 10000000),
		Currency: randomedata.RandomCurrency(),
	}
}

func TestCreateAccount(t *testing.T) {
	SetupTest()
	defer TeardownTest()

	accountArgs := getRandomAcc()
	account, err := accountService.CreateAccount(accountArgs)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, accountArgs.Owner, account.Owner)
	require.Equal(t, accountArgs.Balance, account.Balance)
	require.Equal(t, accountArgs.Currency, account.Currency)
	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccountById(t *testing.T) {
	SetupTest()
	defer TeardownTest()

	accountArgs := getRandomAcc()
	createdAcc, err := accountService.CreateAccount(accountArgs)

	require.NoError(t, err)
	require.NotEmpty(t, createdAcc)

	account, err := accountService.GetAccountById(createdAcc.Id)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, accountArgs.Owner, account.Owner)
	require.Equal(t, accountArgs.Balance, account.Balance)
	require.Equal(t, accountArgs.Currency, account.Currency)
	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreatedAt)
}
