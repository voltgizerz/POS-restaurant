package service

import (
	"os"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/app/mocks"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
	"go.uber.org/mock/gomock"
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
}

func NewMock(t *testing.T) (*gomock.Controller, *MockObject) {
	setTestENV()
	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockAuthService := mocks.NewMockIAuth(ctrl)

	mockObj := &MockObject{
		MockUserRepo:    mockUserRepo,
		MockAuthService: mockAuthService,
	}

	return ctrl, mockObj
}

func setTestENV() {
	os.Setenv("GO_ENV", "unit_test")
}
