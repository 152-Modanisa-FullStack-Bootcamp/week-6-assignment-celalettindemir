// Code generated by MockGen. DO NOT EDIT.
// Source: service/wallet_service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	model "wallet/model"

	gomock "github.com/golang/mock/gomock"
)

// MockWalletService is a mock of WalletService interface.
type MockWalletService struct {
	ctrl     *gomock.Controller
	recorder *MockWalletServiceMockRecorder
}

// MockWalletServiceMockRecorder is the mock recorder for MockWalletService.
type MockWalletServiceMockRecorder struct {
	mock *MockWalletService
}

// NewMockWalletService creates a new mock instance.
func NewMockWalletService(ctrl *gomock.Controller) *MockWalletService {
	mock := &MockWalletService{ctrl: ctrl}
	mock.recorder = &MockWalletServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletService) EXPECT() *MockWalletServiceMockRecorder {
	return m.recorder
}

// CashOperation mocks base method.
func (m *MockWalletService) CashOperation(username string, Amount int) (model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CashOperation", username, Amount)
	ret0, _ := ret[0].(model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CashOperation indicates an expected call of CashOperation.
func (mr *MockWalletServiceMockRecorder) CashOperation(username, Amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CashOperation", reflect.TypeOf((*MockWalletService)(nil).CashOperation), username, Amount)
}

// CreateWallet mocks base method.
func (m *MockWalletService) CreateWallet(username string) (model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", username)
	ret0, _ := ret[0].(model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletServiceMockRecorder) CreateWallet(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletService)(nil).CreateWallet), username)
}

// Wallet mocks base method.
func (m *MockWalletService) Wallet(username string) (model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wallet", username)
	ret0, _ := ret[0].(model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wallet indicates an expected call of Wallet.
func (mr *MockWalletServiceMockRecorder) Wallet(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wallet", reflect.TypeOf((*MockWalletService)(nil).Wallet), username)
}

// Wallets mocks base method.
func (m *MockWalletService) Wallets() ([]model.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wallets")
	ret0, _ := ret[0].([]model.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wallets indicates an expected call of Wallets.
func (mr *MockWalletServiceMockRecorder) Wallets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wallets", reflect.TypeOf((*MockWalletService)(nil).Wallets))
}
