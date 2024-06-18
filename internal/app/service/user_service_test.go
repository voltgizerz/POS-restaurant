package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"go.uber.org/mock/gomock"
)

func TestUserService_Login(t *testing.T) {
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
				password: "test-password",
			},
			want: &entity.LoginResponse{
				UserID: 1,
				RoleID: 1,
				Token:  "mock-token",
			},
			wantErr: false,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsernameAndPassword(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entity.UserORM{ID: 1}, nil).Times(1)

				mockObj.MockAuthService.EXPECT().CreateToken(&entity.UserORM{ID: 1}).
					Return(&entity.CreateTokenResponse{Token: "mock-token"}, nil).Times(1)
			},
		},
		{
			name: "ERROR On - GetUserByUsernameAndPassword",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsernameAndPassword(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errors.New("some errors")).Times(1)
			},
		},
		{
			name: "ERROR On - CreateToken",
			args: args{
				ctx:      context.Background(),
				username: "test-user",
				password: "test-password",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockObj *MockObject) {
				mockObj.MockUserRepo.EXPECT().GetUserByUsernameAndPassword(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&entity.UserORM{ID: 1}, nil).Times(1)

				mockObj.MockAuthService.EXPECT().CreateToken(&entity.UserORM{ID: 1}).
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

			service := &UserService{
				authService:    mockObj.MockAuthService,
				userRepository: mockObj.MockUserRepo,
			}

			got, err := service.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Login() = %v, want %v", got, tt.want)
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
					Return(&entity.UserORM{}, nil).Times(1)
				mockObj.MockUserRepo.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).
					Return(int64(1), nil).Times(1)
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

			service := &UserService{
				userRepository: mockObj.MockUserRepo,
			}

			user_data := &entity.User{
				Username: tt.args.username,
				Password: tt.args.password,
				Email:    tt.args.email,
				Name:     tt.args.name,
			}

			got, err := service.Register(tt.args.ctx, *user_data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
