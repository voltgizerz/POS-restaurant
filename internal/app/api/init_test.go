package api

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
	MockAuthHandler *mocks.MockIAuthHandler
}

func NewMock(t *testing.T) (*gomock.Controller, *MockObject) {
	setTestENV()
	ctrl := gomock.NewController(t)
	mockAuthHandler := mocks.NewMockIAuthHandler(ctrl)

	mockObj := &MockObject{
		MockAuthHandler: mockAuthHandler,
	}

	return ctrl, mockObj
}

func setTestENV() {
	os.Setenv("GO_ENV", "unit_test")
}
