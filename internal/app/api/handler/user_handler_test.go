package handler

import (
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestUserHandler_Login(t *testing.T) {
	type args struct {
		c fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		setup   func(mockObj *MockObject)
	}{{}}
	for _, tt := range tests {
		ctrl, mockObj := NewMock(t)
		if tt.setup != nil {
			tt.setup(mockObj)
		}
		defer ctrl.Finish()

		userHandler := &UserHandler{
			authService: mockObj.MockAuthService,
			userService: mockObj.MockUserService,
		}

		t.Run(tt.name, func(t *testing.T) {
			if err := userHandler.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
