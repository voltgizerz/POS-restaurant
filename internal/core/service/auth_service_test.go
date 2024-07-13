package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
)

func TestUserService_Login(t *testing.T) {
	mockUserORM := &models.UserORM{ID: 1, PasswordHashed: "$2a$14$aRI5bAYlMR7jvM2XH/EB1u9cHMpbuNX6FUsLGPnkdWNeN96OCbw0q"}

	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.LoginResponse
		wantErr bool
		setup   func(mockObj *MockObject)
	}{
		{
			name: "SUCCESS - Login",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "felix",
			},
			want: &entity.LoginResponse{
				UserID:    1,
				Token:     "MOCKING-TOKEN",
				TokenType: "Bearer",
			},
			wantErr: false,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).
					Return(mockUserORM, nil).Times(1)

				mockObj.MockJWTService.EXPECT().CreateToken(gomock.Any(), gomock.Any()).
					Return(&entity.CreateTokenResponse{Token: "MOCKING-TOKEN", TokenType: "Bearer"}, nil).Times(1)
			},
		},
		{
			name: "ERROR - on GetUserByUsername",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("some errors")).Times(1)
			},
		},
		{
			name: "ERROR - on VerifyPassword",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).
					Return(&models.UserORM{ID: 1, PasswordHashed: "aasd"}, nil).Times(1)
			},
		},
		{
			name: "ERROR - on CreateToken",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "felix",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).
					Return(mockUserORM, nil).Times(1)

				mockObj.MockJWTService.EXPECT().CreateToken(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("some errors")).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, mockObj := NewMock(t)
			if tt.setup != nil {
				tt.setup(mockObj)
			}
			defer ctrl.Finish()

			service := &AuthService{
				jwtService:     mockObj.MockJWTService,
				txRepository:   mockObj.MockTxRepo,
				userRepository: mockObj.MockUserRepo,
			}

			got, err := service.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Register(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
		name     string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
		setup   func(mockObj *MockObject)
	}{
		{
			name: "SUCCESS - Register",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
				email:    "test-email@email.com",
				name:     "test-name",
			},
			want:    1,
			wantErr: false,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).
					Return(&models.UserORM{}, nil).Times(1)
				mockObj.MockUserRepo.EXPECT().RegisterUser(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(1), nil).Times(1)
			},
		},
		{
			name: "ERROR - GetUserByEmail",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
				email:    "test-email@email.com",
				name:     "test-name",
			},
			want:    0,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).
					Return(&models.UserORM{Username: ""}, errors.New("some error")).Times(1)
			},
		},
		{
			name: "ERROR - Register",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
				email:    "test-email@email.com",
				name:     "test-name",
			},
			want:    0,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).
					Return(&models.UserORM{Username: ""}, errors.New("some error")).Times(1)
				mockObj.MockUserRepo.EXPECT().RegisterUser(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(0), errors.New("some error")).AnyTimes()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, mockObj := NewMock(t)
			if tt.setup != nil {
				tt.setup(mockObj)
			}
			defer ctrl.Finish()

			service := &AuthService{
				userRepository: mockObj.MockUserRepo,
			}

			userData := &entity.User{
				Username: tt.args.username,
				Password: tt.args.password,
				Email:    tt.args.email,
				Name:     tt.args.name,
			}

			got, err := service.Register(tt.args.ctx, *userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthService(t *testing.T) {
	type args struct {
		i interactor.AuthService
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		{
			name: "SUCCESS",
			args: args{i: interactor.AuthService{}},
			want: &AuthService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}
