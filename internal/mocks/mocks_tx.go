// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/tx_ports.go
//
// Generated by this command:
//
//	mockgen -source=./internal/core/ports/tx_ports.go -destination=./internal/mocks/mocks_tx.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockITxRepository is a mock of ITxRepository interface.
type MockITxRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITxRepositoryMockRecorder
}

// MockITxRepositoryMockRecorder is the mock recorder for MockITxRepository.
type MockITxRepositoryMockRecorder struct {
	mock *MockITxRepository
}

// NewMockITxRepository creates a new mock instance.
func NewMockITxRepository(ctrl *gomock.Controller) *MockITxRepository {
	mock := &MockITxRepository{ctrl: ctrl}
	mock.recorder = &MockITxRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITxRepository) EXPECT() *MockITxRepositoryMockRecorder {
	return m.recorder
}

// CommitTransaction mocks base method.
func (m *MockITxRepository) CommitTransaction(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockITxRepositoryMockRecorder) CommitTransaction(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockITxRepository)(nil).CommitTransaction), ctx, tx)
}

// RollbackTransaction mocks base method.
func (m *MockITxRepository) RollbackTransaction(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTransaction", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockITxRepositoryMockRecorder) RollbackTransaction(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockITxRepository)(nil).RollbackTransaction), ctx, tx)
}

// StartTransaction mocks base method.
func (m *MockITxRepository) StartTransaction(ctx context.Context) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTransaction", ctx)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTransaction indicates an expected call of StartTransaction.
func (mr *MockITxRepositoryMockRecorder) StartTransaction(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTransaction", reflect.TypeOf((*MockITxRepository)(nil).StartTransaction), ctx)
}
