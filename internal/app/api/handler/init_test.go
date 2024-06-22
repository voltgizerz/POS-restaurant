package handler

import (
	"os"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/voltgizerz/POS-restaurant/internal/app/mocks"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

func TestMain(m *testing.M) {
	// Setup code (if any)
	// ...
	logger.Init()
	// Run the tests
	m.Run()
}

type MockObject struct {
	MockUserRepo    *mocks.MockIUserRepository
	MockAuthService *mocks.MockIAuth
	MockUserService *mocks.MockIUserService
}

func NewMock(t *testing.T) (*gomock.Controller, *MockObject) {
	setTestENV()
	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockAuthService := mocks.NewMockIAuth(ctrl)
	mockUserService := mocks.NewMockIUserService(ctrl)

	mockObj := &MockObject{
		MockUserRepo:    mockUserRepo,
		MockAuthService: mockAuthService,
		MockUserService: mockUserService,
	}

	return ctrl, mockObj
}

func setTestENV() {
	os.Setenv("GO_ENV", "unit_test")
}
