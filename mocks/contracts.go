// Code generated by MockGen. DO NOT EDIT.
// Source: ./contracts/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "bitbucket.org/artie_inc/contract-service/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockContractTransactionHandler is a mock of ContractTransactionHandler interface.
type MockContractTransactionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockContractTransactionHandlerMockRecorder
}

// MockContractTransactionHandlerMockRecorder is the mock recorder for MockContractTransactionHandler.
type MockContractTransactionHandlerMockRecorder struct {
	mock *MockContractTransactionHandler
}

// NewMockContractTransactionHandler creates a new mock instance.
func NewMockContractTransactionHandler(ctrl *gomock.Controller) *MockContractTransactionHandler {
	mock := &MockContractTransactionHandler{ctrl: ctrl}
	mock.recorder = &MockContractTransactionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractTransactionHandler) EXPECT() *MockContractTransactionHandlerMockRecorder {
	return m.recorder
}

// BuildTransaction mocks base method.
func (m *MockContractTransactionHandler) BuildTransaction(ctx context.Context, senderInHash bool, msgSender, functionName string, arguments [][]byte, value string, contract *storage.Contract) (*storage.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTransaction", ctx, senderInHash, msgSender, functionName, arguments, value, contract)
	ret0, _ := ret[0].(*storage.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTransaction indicates an expected call of BuildTransaction.
func (mr *MockContractTransactionHandlerMockRecorder) BuildTransaction(ctx, senderInHash, msgSender, functionName, arguments, value, contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTransaction", reflect.TypeOf((*MockContractTransactionHandler)(nil).BuildTransaction), ctx, senderInHash, msgSender, functionName, arguments, value, contract)
}

// CompleteTransaction mocks base method.
func (m *MockContractTransactionHandler) CompleteTransaction(ctx context.Context, address, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTransaction", ctx, address, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteTransaction indicates an expected call of CompleteTransaction.
func (mr *MockContractTransactionHandlerMockRecorder) CompleteTransaction(ctx, address, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTransaction", reflect.TypeOf((*MockContractTransactionHandler)(nil).CompleteTransaction), ctx, address, hash)
}

// DeleteTransaction mocks base method.
func (m *MockContractTransactionHandler) DeleteTransaction(ctx context.Context, address, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", ctx, address, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockContractTransactionHandlerMockRecorder) DeleteTransaction(ctx, address, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockContractTransactionHandler)(nil).DeleteTransaction), ctx, address, hash)
}

// GetAllTransactions mocks base method.
func (m *MockContractTransactionHandler) GetAllTransactions(ctx context.Context, address string) ([]*storage.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransactions", ctx, address)
	ret0, _ := ret[0].([]*storage.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransactions indicates an expected call of GetAllTransactions.
func (mr *MockContractTransactionHandlerMockRecorder) GetAllTransactions(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransactions", reflect.TypeOf((*MockContractTransactionHandler)(nil).GetAllTransactions), ctx, address)
}

// GetContract mocks base method.
func (m *MockContractTransactionHandler) GetContract(ctx context.Context, address string) (*storage.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", ctx, address)
	ret0, _ := ret[0].(*storage.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockContractTransactionHandlerMockRecorder) GetContract(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractTransactionHandler)(nil).GetContract), ctx, address)
}

// GetTransactions mocks base method.
func (m *MockContractTransactionHandler) GetTransactions(ctx context.Context, address string) ([]*storage.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", ctx, address)
	ret0, _ := ret[0].([]*storage.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockContractTransactionHandlerMockRecorder) GetTransactions(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockContractTransactionHandler)(nil).GetTransactions), ctx, address)
}

// StoreTransaction mocks base method.
func (m *MockContractTransactionHandler) StoreTransaction(ctx context.Context, token *storage.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreTransaction", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreTransaction indicates an expected call of StoreTransaction.
func (mr *MockContractTransactionHandlerMockRecorder) StoreTransaction(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreTransaction", reflect.TypeOf((*MockContractTransactionHandler)(nil).StoreTransaction), ctx, token)
}

// MockContractManagerHandler is a mock of ContractManagerHandler interface.
type MockContractManagerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagerHandlerMockRecorder
}

// MockContractManagerHandlerMockRecorder is the mock recorder for MockContractManagerHandler.
type MockContractManagerHandlerMockRecorder struct {
	mock *MockContractManagerHandler
}

// NewMockContractManagerHandler creates a new mock instance.
func NewMockContractManagerHandler(ctrl *gomock.Controller) *MockContractManagerHandler {
	mock := &MockContractManagerHandler{ctrl: ctrl}
	mock.recorder = &MockContractManagerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagerHandler) EXPECT() *MockContractManagerHandlerMockRecorder {
	return m.recorder
}

// DeleteContract mocks base method.
func (m *MockContractManagerHandler) DeleteContract(ctx context.Context, address, owner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContract", ctx, address, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContract indicates an expected call of DeleteContract.
func (mr *MockContractManagerHandlerMockRecorder) DeleteContract(ctx, address, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContract", reflect.TypeOf((*MockContractManagerHandler)(nil).DeleteContract), ctx, address, owner)
}

// GetContract mocks base method.
func (m *MockContractManagerHandler) GetContract(ctx context.Context, address string) (*storage.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", ctx, address)
	ret0, _ := ret[0].(*storage.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockContractManagerHandlerMockRecorder) GetContract(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractManagerHandler)(nil).GetContract), ctx, address)
}

// ListContracts mocks base method.
func (m *MockContractManagerHandler) ListContracts(ctx context.Context, owner string) ([]*storage.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContracts", ctx, owner)
	ret0, _ := ret[0].([]*storage.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContracts indicates an expected call of ListContracts.
func (mr *MockContractManagerHandlerMockRecorder) ListContracts(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContracts", reflect.TypeOf((*MockContractManagerHandler)(nil).ListContracts), ctx, owner)
}

// StoreContract mocks base method.
func (m *MockContractManagerHandler) StoreContract(ctx context.Context, contract *storage.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreContract", ctx, contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreContract indicates an expected call of StoreContract.
func (mr *MockContractManagerHandlerMockRecorder) StoreContract(ctx, contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreContract", reflect.TypeOf((*MockContractManagerHandler)(nil).StoreContract), ctx, contract)
}
