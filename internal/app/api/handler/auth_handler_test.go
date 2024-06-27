package handler

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"

	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

func TestAuthHandler_Login(t *testing.T) {
	app := fiber.New()

	tests := []struct {
		name    string
		h       *AuthHandler
		wantErr bool
		setup   func(mockObj *MockObject) fiber.Ctx
	}{
		{
			name:    "SUCCESS - Login",
			wantErr: false,
			setup: func(mockObj *MockObject) fiber.Ctx {
				mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
				mockCtx.Locals(constants.CTXKeyRequestID, "mock-req-id")

				data, _ := sonic.Marshal(entity.LoginRequest{
					Username: "fELIX",
					Password: "fELIX",
				})
				mockCtx.Request().SetBodyString(string(data))

				// mockObj.MockAuthService.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).
				// 	Return(&entity.LoginResponse{}, nil).Times(1)
				return mockCtx
			},
		},
	}
	for _, tt := range tests {
		ctrl, mockObj := NewMock(t)
		defer ctrl.Finish()

		authHandler := &AuthHandler{
			authService: mockObj.MockAuthService,
		}

		c := tt.setup(mockObj)

		t.Run(tt.name, func(t *testing.T) {
			if err := authHandler.Login(c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
