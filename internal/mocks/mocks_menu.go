// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/ports/menu_ports.go
//
// Generated by this command:
//
//	mockgen -source=./internal/app/ports/menu_ports.go -destination=./internal/mocks/mocks_menu.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v3"
	entity "github.com/voltgizerz/POS-restaurant/internal/core/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockIMenuHandler is a mock of IMenuHandler interface.
type MockIMenuHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIMenuHandlerMockRecorder
}

// MockIMenuHandlerMockRecorder is the mock recorder for MockIMenuHandler.
type MockIMenuHandlerMockRecorder struct {
	mock *MockIMenuHandler
}

// NewMockIMenuHandler creates a new mock instance.
func NewMockIMenuHandler(ctrl *gomock.Controller) *MockIMenuHandler {
	mock := &MockIMenuHandler{ctrl: ctrl}
	mock.recorder = &MockIMenuHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMenuHandler) EXPECT() *MockIMenuHandlerMockRecorder {
	return m.recorder
}

// AddMenu mocks base method.
func (m *MockIMenuHandler) AddMenu(c fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMenu", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMenu indicates an expected call of AddMenu.
func (mr *MockIMenuHandlerMockRecorder) AddMenu(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMenu", reflect.TypeOf((*MockIMenuHandler)(nil).AddMenu), c)
}

// GetMenuByUserID mocks base method.
func (m *MockIMenuHandler) GetMenuByUserID(c fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuByUserID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMenuByUserID indicates an expected call of GetMenuByUserID.
func (mr *MockIMenuHandlerMockRecorder) GetMenuByUserID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuByUserID", reflect.TypeOf((*MockIMenuHandler)(nil).GetMenuByUserID), c)
}

// UpdateActiveMenuBatchByUserID mocks base method.
func (m *MockIMenuHandler) UpdateActiveMenuBatchByUserID(c fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuBatchByUserID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActiveMenuBatchByUserID indicates an expected call of UpdateActiveMenuBatchByUserID.
func (mr *MockIMenuHandlerMockRecorder) UpdateActiveMenuBatchByUserID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuBatchByUserID", reflect.TypeOf((*MockIMenuHandler)(nil).UpdateActiveMenuBatchByUserID), c)
}

// UpdateActiveMenuByMenuID mocks base method.
func (m *MockIMenuHandler) UpdateActiveMenuByMenuID(c fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuByMenuID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActiveMenuByMenuID indicates an expected call of UpdateActiveMenuByMenuID.
func (mr *MockIMenuHandlerMockRecorder) UpdateActiveMenuByMenuID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuByMenuID", reflect.TypeOf((*MockIMenuHandler)(nil).UpdateActiveMenuByMenuID), c)
}

// UpdateMenuByMenuID mocks base method.
func (m *MockIMenuHandler) UpdateMenuByMenuID(c fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMenuByMenuID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMenuByMenuID indicates an expected call of UpdateMenuByMenuID.
func (mr *MockIMenuHandlerMockRecorder) UpdateMenuByMenuID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuByMenuID", reflect.TypeOf((*MockIMenuHandler)(nil).UpdateMenuByMenuID), c)
}

// MockIMenuService is a mock of IMenuService interface.
type MockIMenuService struct {
	ctrl     *gomock.Controller
	recorder *MockIMenuServiceMockRecorder
}

// MockIMenuServiceMockRecorder is the mock recorder for MockIMenuService.
type MockIMenuServiceMockRecorder struct {
	mock *MockIMenuService
}

// NewMockIMenuService creates a new mock instance.
func NewMockIMenuService(ctrl *gomock.Controller) *MockIMenuService {
	mock := &MockIMenuService{ctrl: ctrl}
	mock.recorder = &MockIMenuServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMenuService) EXPECT() *MockIMenuServiceMockRecorder {
	return m.recorder
}

// GetMenu mocks base method.
func (m *MockIMenuService) GetMenu(ctx context.Context, menuID int64) ([]*entity.MenuResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", ctx, menuID)
	ret0, _ := ret[0].([]*entity.MenuResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockIMenuServiceMockRecorder) GetMenu(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockIMenuService)(nil).GetMenu), ctx, menuID)
}

// RegisterMenu mocks base method.
func (m *MockIMenuService) RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMenu", ctx, menuData)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterMenu indicates an expected call of RegisterMenu.
func (mr *MockIMenuServiceMockRecorder) RegisterMenu(ctx, menuData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMenu", reflect.TypeOf((*MockIMenuService)(nil).RegisterMenu), ctx, menuData)
}

// UpdateActiveMenuBatchUserID mocks base method.
func (m *MockIMenuService) UpdateActiveMenuBatchUserID(ctx context.Context, userID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuBatchUserID", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActiveMenuBatchUserID indicates an expected call of UpdateActiveMenuBatchUserID.
func (mr *MockIMenuServiceMockRecorder) UpdateActiveMenuBatchUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuBatchUserID", reflect.TypeOf((*MockIMenuService)(nil).UpdateActiveMenuBatchUserID), ctx, userID)
}

// UpdateActiveMenuID mocks base method.
func (m *MockIMenuService) UpdateActiveMenuID(ctx context.Context, menuID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuID", ctx, menuID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActiveMenuID indicates an expected call of UpdateActiveMenuID.
func (mr *MockIMenuServiceMockRecorder) UpdateActiveMenuID(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuID", reflect.TypeOf((*MockIMenuService)(nil).UpdateActiveMenuID), ctx, menuID)
}

// UpdateMenuID mocks base method.
func (m *MockIMenuService) UpdateMenuID(ctx context.Context, menuData entity.Menu) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMenuID", ctx, menuData)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMenuID indicates an expected call of UpdateMenuID.
func (mr *MockIMenuServiceMockRecorder) UpdateMenuID(ctx, menuData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuID", reflect.TypeOf((*MockIMenuService)(nil).UpdateMenuID), ctx, menuData)
}

// MockIMenuRepository is a mock of IMenuRepository interface.
type MockIMenuRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMenuRepositoryMockRecorder
}

// MockIMenuRepositoryMockRecorder is the mock recorder for MockIMenuRepository.
type MockIMenuRepositoryMockRecorder struct {
	mock *MockIMenuRepository
}

// NewMockIMenuRepository creates a new mock instance.
func NewMockIMenuRepository(ctrl *gomock.Controller) *MockIMenuRepository {
	mock := &MockIMenuRepository{ctrl: ctrl}
	mock.recorder = &MockIMenuRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMenuRepository) EXPECT() *MockIMenuRepositoryMockRecorder {
	return m.recorder
}

// AddMenu mocks base method.
func (m *MockIMenuRepository) AddMenu(ctx context.Context, menuData entity.MenuORM) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMenu", ctx, menuData)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMenu indicates an expected call of AddMenu.
func (mr *MockIMenuRepositoryMockRecorder) AddMenu(ctx, menuData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMenu", reflect.TypeOf((*MockIMenuRepository)(nil).AddMenu), ctx, menuData)
}

// FetchMenuById mocks base method.
func (m *MockIMenuRepository) FetchMenuById(ctx context.Context, userID int64) ([]*entity.MenuORM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMenuById", ctx, userID)
	ret0, _ := ret[0].([]*entity.MenuORM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMenuById indicates an expected call of FetchMenuById.
func (mr *MockIMenuRepositoryMockRecorder) FetchMenuById(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMenuById", reflect.TypeOf((*MockIMenuRepository)(nil).FetchMenuById), ctx, userID)
}

// UpdateActiveMenuBatchUser mocks base method.
func (m *MockIMenuRepository) UpdateActiveMenuBatchUser(ctx context.Context, userID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuBatchUser", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActiveMenuBatchUser indicates an expected call of UpdateActiveMenuBatchUser.
func (mr *MockIMenuRepositoryMockRecorder) UpdateActiveMenuBatchUser(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuBatchUser", reflect.TypeOf((*MockIMenuRepository)(nil).UpdateActiveMenuBatchUser), ctx, userID)
}

// UpdateActiveMenuByMenuID mocks base method.
func (m *MockIMenuRepository) UpdateActiveMenuByMenuID(ctx context.Context, menuID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActiveMenuByMenuID", ctx, menuID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActiveMenuByMenuID indicates an expected call of UpdateActiveMenuByMenuID.
func (mr *MockIMenuRepositoryMockRecorder) UpdateActiveMenuByMenuID(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActiveMenuByMenuID", reflect.TypeOf((*MockIMenuRepository)(nil).UpdateActiveMenuByMenuID), ctx, menuID)
}

// UpdateMenuByMenuID mocks base method.
func (m *MockIMenuRepository) UpdateMenuByMenuID(ctx context.Context, menuData *entity.MenuORM) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMenuByMenuID", ctx, menuData)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMenuByMenuID indicates an expected call of UpdateMenuByMenuID.
func (mr *MockIMenuRepositoryMockRecorder) UpdateMenuByMenuID(ctx, menuData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMenuByMenuID", reflect.TypeOf((*MockIMenuRepository)(nil).UpdateMenuByMenuID), ctx, menuData)
}
