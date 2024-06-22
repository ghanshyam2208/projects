package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"banking_app2/cmd/internals/dto"
	mock_services "banking_app2/mock"
)

var (
	ctrl        *gomock.Controller
	mockService *mock_services.MockIAccountService
	handler     AccountHandlers
)

func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockService = mock_services.NewMockIAccountService(ctrl)
	handler = AccountHandlers{service: mockService}
}

func teardown() {
	ctrl.Finish()
}

func SetupTest(t *testing.T) {
	setup(t)
}

func TeardownTest(t *testing.T) {
	teardown()
}

func TestAccountHandlers_GetAllAccounts(t *testing.T) {
	SetupTest(t)
	defer TeardownTest(t)

	// Expected data from mock service
	expectedAccounts := []dto.AccountDto{
		{Id: 1, Owner: "John"},
		{Id: 2, Owner: "Jane"},
	}

	// Set up expectations on the mock service
	mockService.EXPECT().GetAllAccounts(1, 10).Return(expectedAccounts, nil)

	// Echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/accounts?page=1&pageSize=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Handler call
	err := handler.GetAllAccounts(c)
	if err != nil {
		t.Fatalf("handler error: %v", err)
	}

	// Check response status code
	assert.Equal(t, http.StatusOK, rec.Code, "status code does not match expected")

	// Decode response body
	var response map[string]interface{}
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	// Check accounts field
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("unexpected type for Data field: %T", response["Data"])
	}

	accounts, ok := data["accounts"].([]interface{})
	assert.True(t, ok, "accounts should be a []interface{}")
	assert.Equal(t, len(expectedAccounts), len(accounts), "number of accounts should match")
}
