// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/adapters/ports/user_ports.go
//
// Generated by this command:
//
//	mockgen -source=./internal/adapters/ports/user_ports.go -destination=./internal/mocks/mocks_user.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/voltgizerz/POS-restaurant/internal/core/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockIUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserORM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.UserORM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserRepositoryMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserByUsername mocks base method.
func (m *MockIUserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.UserORM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(*entity.UserORM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockIUserRepositoryMockRecorder) GetUserByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByUsername), ctx, username)
}

// GetUserByUsernameAndPassword mocks base method.
func (m *MockIUserRepository) GetUserByUsernameAndPassword(ctx context.Context, username, hashPassword string) (*entity.UserORM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsernameAndPassword", ctx, username, hashPassword)
	ret0, _ := ret[0].(*entity.UserORM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsernameAndPassword indicates an expected call of GetUserByUsernameAndPassword.
func (mr *MockIUserRepositoryMockRecorder) GetUserByUsernameAndPassword(ctx, username, hashPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsernameAndPassword", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByUsernameAndPassword), ctx, username, hashPassword)
}

// RegisterUser mocks base method.
func (m *MockIUserRepository) RegisterUser(ctx context.Context, userData entity.UserORM) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, userData)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserRepositoryMockRecorder) RegisterUser(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserRepository)(nil).RegisterUser), ctx, userData)
}