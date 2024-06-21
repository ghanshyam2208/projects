package services

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	"fmt"
	"os"
	"testing"
)

var accountService *AccountService

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
	// Initialize the service with a repository
	accountService = NewAccountService(repositories.NewAccountsRepo())
}

// teardown cleans up resources after all tests
func teardown() {
	// Perform any necessary cleanup
}

// SetupTest initializes resources before each test
func SetupTest() {
	// Clean the database or repository before each test
	// repositories.ClearAllData()
}

// TeardownTest cleans up resources after each test
func TeardownTest() {
	// Optionally clean up after each test if needed
}

func TestCreateAccount(t *testing.T) {
	SetupTest()
	defer TeardownTest()

	args := dto.CreateAccountDto{
		Owner:    "sample user",
		Balance:  1000,
		Currency: "INR",
	}

	account, err := accountService.CreateAccount(args)

	if err != nil {
		t.Errorf("Error creating account: %v", err)
	}
	fmt.Println(account)
}
