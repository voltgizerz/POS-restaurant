package handler

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"go.uber.org/mock/gomock"
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
				mockCtx.Request().Header.Set("Content-Type", "application/json")

				data, _ := sonic.Marshal(entity.LoginRequest{
					Username: "fELIX",
					Password: "fELIX",
				})
				mockCtx.Request().SetBodyString(string(data))

				mockObj.MockAuthService.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entity.LoginResponse{}, nil).Times(1)
				return mockCtx
			},
		},
		{
			name:    "ERROR - on Bind",
			wantErr: false,
			setup: func(mockObj *MockObject) fiber.Ctx {
				mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
				mockCtx.Locals(constants.CTXKeyRequestID, "mock-req-id")
				mockCtx.Request().Header.Set("Content-Type", "random")

				return mockCtx
			},
		},
		{
			name:    "Error - on Validator",
			wantErr: false,
			setup: func(mockObj *MockObject) fiber.Ctx {
				mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
				mockCtx.Locals(constants.CTXKeyRequestID, "mock-req-id")
				mockCtx.Request().Header.Set("Content-Type", "application/json")

				data, _ := sonic.Marshal(entity.LoginRequest{
					Username: "",
					Password: "fELIX",
				})
				mockCtx.Request().SetBodyString(string(data))

				return mockCtx
			},
		},
		{
			name:    "Error - on Login",
			wantErr: false,
			setup: func(mockObj *MockObject) fiber.Ctx {
				mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
				mockCtx.Locals(constants.CTXKeyRequestID, "mock-req-id")
				mockCtx.Request().Header.Set("Content-Type", "application/json")

				data, _ := sonic.Marshal(entity.LoginRequest{
					Username: "felix",
					Password: "fELIX",
				})
				mockCtx.Request().SetBodyString(string(data))

				mockObj.MockAuthService.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("some errors")).Times(1)
				return mockCtx
			},
		},
		{
			name:    "Error - on Login got no rows",
			wantErr: false,
			setup: func(mockObj *MockObject) fiber.Ctx {
				mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
				mockCtx.Locals(constants.CTXKeyRequestID, "mock-req-id")
				mockCtx.Request().Header.Set("Content-Type", "application/json")

				data, _ := sonic.Marshal(entity.LoginRequest{
					Username: "felix",
					Password: "fELIX",
				})
				mockCtx.Request().SetBodyString(string(data))

				mockObj.MockAuthService.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, sql.ErrNoRows).Times(1)
				return mockCtx
			},
		},
	}
	for _, tt := range tests {
		ctrl, mockObj := NewMock(t)
		defer ctrl.Finish()

		c := tt.setup(mockObj)
		authHandler := &AuthHandler{
			authService: mockObj.MockAuthService,
		}

		t.Run(tt.name, func(t *testing.T) {
			if err := authHandler.Login(c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		i interactor.AuthHandler
	}
	tests := []struct {
		name string
		args args
		want *AuthHandler
	}{
		{
			name: "SUCCESS",
			args: args{},
			want: &AuthHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHandler(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
