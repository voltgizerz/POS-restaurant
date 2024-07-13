package service

import (
	"os"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/mocks"
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
	MockJWTService *mocks.MockIJWTAuth
	MockUserRepo   *mocks.MockIUserRepository
	MockTxRepo     *mocks.MockITxRepository
}

func NewMock(t *testing.T) (*gomock.Controller, *MockObject) {
	setTestENV()
	ctrl := gomock.NewController(t)
	mockJWTService := mocks.NewMockIJWTAuth(ctrl)
	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockTxRepo := mocks.NewMockITxRepository(ctrl)

	mockObj := &MockObject{
		MockJWTService: mockJWTService,
		MockUserRepo:   mockUserRepo,
		MockTxRepo:     mockTxRepo,
	}

	return ctrl, mockObj
}

func setTestENV() {
	os.Setenv("GO_ENV", "unit_test")
}
