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
	MockUserRepo   *mocks.MockIUserRepository
	MockJWTService *mocks.MockIJWTAuth
}

func NewMock(t *testing.T) (*gomock.Controller, *MockObject) {
	setTestENV()
	ctrl := gomock.NewController(t)
	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockJWTService := mocks.NewMockIJWTAuth(ctrl)

	mockObj := &MockObject{
		MockUserRepo:   mockUserRepo,
		MockJWTService: mockJWTService,
	}

	return ctrl, mockObj
}

func setTestENV() {
	os.Setenv("GO_ENV", "unit_test")
}
