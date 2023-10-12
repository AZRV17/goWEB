// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	repository "github.com/AZRV17/goWEB/internal/repository"
	service "github.com/AZRV17/goWEB/internal/service"
	gomock "github.com/golang/mock/gomock"
)

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccount) CreateAccount(input service.CreateAccountInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountMockRecorder) CreateAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccount)(nil).CreateAccount), input)
}

// DeleteAccount mocks base method.
func (m *MockAccount) DeleteAccount(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockAccountMockRecorder) DeleteAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccount)(nil).DeleteAccount), id)
}

// GetAccount mocks base method.
func (m *MockAccount) GetAccount(id int64) (repository.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", id)
	ret0, _ := ret[0].(repository.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountMockRecorder) GetAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccount)(nil).GetAccount), id)
}

// GetAllAccounts mocks base method.
func (m *MockAccount) GetAllAccounts() ([]repository.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccounts")
	ret0, _ := ret[0].([]repository.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccounts indicates an expected call of GetAllAccounts.
func (mr *MockAccountMockRecorder) GetAllAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccounts", reflect.TypeOf((*MockAccount)(nil).GetAllAccounts))
}

// UpdateAccount mocks base method.
func (m *MockAccount) UpdateAccount(input service.UpdateAccountInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountMockRecorder) UpdateAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccount)(nil).UpdateAccount), input)
}

// MockEntry is a mock of Entry interface.
type MockEntry struct {
	ctrl     *gomock.Controller
	recorder *MockEntryMockRecorder
}

// MockEntryMockRecorder is the mock recorder for MockEntry.
type MockEntryMockRecorder struct {
	mock *MockEntry
}

// NewMockEntry creates a new mock instance.
func NewMockEntry(ctrl *gomock.Controller) *MockEntry {
	mock := &MockEntry{ctrl: ctrl}
	mock.recorder = &MockEntryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntry) EXPECT() *MockEntryMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method.
func (m *MockEntry) CreateEntry(input service.CreateEntryInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockEntryMockRecorder) CreateEntry(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockEntry)(nil).CreateEntry), input)
}

// DeleteEntry mocks base method.
func (m *MockEntry) DeleteEntry(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry.
func (mr *MockEntryMockRecorder) DeleteEntry(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockEntry)(nil).DeleteEntry), id)
}

// GetAllEntries mocks base method.
func (m *MockEntry) GetAllEntries() ([]repository.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEntries")
	ret0, _ := ret[0].([]repository.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEntries indicates an expected call of GetAllEntries.
func (mr *MockEntryMockRecorder) GetAllEntries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEntries", reflect.TypeOf((*MockEntry)(nil).GetAllEntries))
}

// GetEntry mocks base method.
func (m *MockEntry) GetEntry(id int64) (repository.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", id)
	ret0, _ := ret[0].(repository.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockEntryMockRecorder) GetEntry(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockEntry)(nil).GetEntry), id)
}

// UpdateEntry mocks base method.
func (m *MockEntry) UpdateEntry(input service.UpdateEntryInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntry", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntry indicates an expected call of UpdateEntry.
func (mr *MockEntryMockRecorder) UpdateEntry(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockEntry)(nil).UpdateEntry), input)
}

// MockTransferServiceInterface is a mock of TransferServiceInterface interface.
type MockTransferServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransferServiceInterfaceMockRecorder
}

// MockTransferServiceInterfaceMockRecorder is the mock recorder for MockTransferServiceInterface.
type MockTransferServiceInterfaceMockRecorder struct {
	mock *MockTransferServiceInterface
}

// NewMockTransferServiceInterface creates a new mock instance.
func NewMockTransferServiceInterface(ctrl *gomock.Controller) *MockTransferServiceInterface {
	mock := &MockTransferServiceInterface{ctrl: ctrl}
	mock.recorder = &MockTransferServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferServiceInterface) EXPECT() *MockTransferServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateTransfer mocks base method.
func (m *MockTransferServiceInterface) CreateTransfer(input service.CreateTransferInput) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", input)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockTransferServiceInterfaceMockRecorder) CreateTransfer(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockTransferServiceInterface)(nil).CreateTransfer), input)
}

// DeleteTransfer mocks base method.
func (m *MockTransferServiceInterface) DeleteTransfer(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfer", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfer indicates an expected call of DeleteTransfer.
func (mr *MockTransferServiceInterfaceMockRecorder) DeleteTransfer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfer", reflect.TypeOf((*MockTransferServiceInterface)(nil).DeleteTransfer), id)
}

// GetAllTransfers mocks base method.
func (m *MockTransferServiceInterface) GetAllTransfers() ([]repository.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransfers")
	ret0, _ := ret[0].([]repository.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransfers indicates an expected call of GetAllTransfers.
func (mr *MockTransferServiceInterfaceMockRecorder) GetAllTransfers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransfers", reflect.TypeOf((*MockTransferServiceInterface)(nil).GetAllTransfers))
}

// GetTransfer mocks base method.
func (m *MockTransferServiceInterface) GetTransfer(id int64) (repository.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", id)
	ret0, _ := ret[0].(repository.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockTransferServiceInterfaceMockRecorder) GetTransfer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockTransferServiceInterface)(nil).GetTransfer), id)
}

// UpdateTransfer mocks base method.
func (m *MockTransferServiceInterface) UpdateTransfer(input service.UpdateTransferInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransfer", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransfer indicates an expected call of UpdateTransfer.
func (mr *MockTransferServiceInterfaceMockRecorder) UpdateTransfer(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransfer", reflect.TypeOf((*MockTransferServiceInterface)(nil).UpdateTransfer), input)
}