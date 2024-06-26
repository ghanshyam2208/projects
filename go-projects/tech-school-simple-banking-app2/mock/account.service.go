// Code generated by MockGen. DO NOT EDIT.
// Source: ./cmd/internals/services/account.service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	dto "banking_app2/cmd/internals/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAccountService is a mock of IAccountService interface.
type MockIAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountServiceMockRecorder
}

// MockIAccountServiceMockRecorder is the mock recorder for MockIAccountService.
type MockIAccountServiceMockRecorder struct {
	mock *MockIAccountService
}

// NewMockIAccountService creates a new mock instance.
func NewMockIAccountService(ctrl *gomock.Controller) *MockIAccountService {
	mock := &MockIAccountService{ctrl: ctrl}
	mock.recorder = &MockIAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountService) EXPECT() *MockIAccountServiceMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockIAccountService) CreateAccount(arg0 dto.CreateAccountDto) (dto.AccountDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(dto.AccountDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockIAccountServiceMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIAccountService)(nil).CreateAccount), arg0)
}

// DeleteAccount mocks base method.
func (m *MockIAccountService) DeleteAccount(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockIAccountServiceMockRecorder) DeleteAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockIAccountService)(nil).DeleteAccount), arg0)
}

// GetAccountById mocks base method.
func (m *MockIAccountService) GetAccountById(arg0 int64) (dto.AccountDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", arg0)
	ret0, _ := ret[0].(dto.AccountDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockIAccountServiceMockRecorder) GetAccountById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockIAccountService)(nil).GetAccountById), arg0)
}

// GetAllAccounts mocks base method.
func (m *MockIAccountService) GetAllAccounts(arg0, arg1 int) ([]dto.AccountDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts", arg0, arg1)
	ret0, _ := ret[0].([]dto.AccountDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockIAccountServiceMockRecorder) GetAllAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockIAccountService)(nil).GetAllAccounts), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockIAccountService) UpdateAccount(arg0 dto.UpdateAccountDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockIAccountServiceMockRecorder) UpdateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockIAccountService)(nil).UpdateAccount), arg0)
}
