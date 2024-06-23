package handler

import (
	"testing"

	"github.com/gofiber/fiber/v3"
)

// func TestUserHandler_Login(t *testing.T) {
// 	app := fiber.New()
// 	mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
// 	mockCtx.Request().SetBodyString(`{"username":"user","password":"password"}`)

// 	type args struct {
// 		c fiber.Ctx
// 	}
// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		wantCode int
// 		setup    func(mockObj *MockObject)
// 	}{{
// 		name: "SUCCESS - Login",
// 		args: args{c: mockCtx},
// 		setup: func(mockObj *MockObject) {
// 		},
// 		wantCode: fiber.StatusOK,
// 	}}
// 	for _, tt := range tests {
// 		ctrl, mockObj := NewMock(t)
// 		if tt.setup != nil {
// 			tt.setup(mockObj)
// 		}
// 		defer ctrl.Finish()

// 		authHandler := &AuthHandler{
// 			userService: mockObj.MockUserService,
// 		}

// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := authHandler.Login(tt.args.c); (err != nil) != tt.wantErr {
// 				t.Errorf("AuthHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			var responseBody successResponse
// 			err := json.Unmarshal(tt.args.c.Response().Body(), &responseBody)
// 			require.NoError(t, err)
// 			logger.Log.Println(responseBody)
// 			assert.Equal(t, responseBody.Code, tt.wantCode)
// 			assert.True(t, gock.IsDone())
// 		})
// 	}
// }

func TestAuthHandler_Login(t *testing.T) {
	type args struct {
		c fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AuthHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHandler.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
