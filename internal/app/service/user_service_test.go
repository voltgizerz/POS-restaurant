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

				mockObj.MockAuthService.EXPECT().CreateToken(gomock.Any(), &entity.UserORM{ID: 1}).
					Return(&entity.CreateTokenResponse{Token: "mock-token"}, nil).Times(1)
			},
		},
		{
			name: "ERROR - on GetUserByUsernameAndPassword",
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
			name: "ERROR - on CreateToken",
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

				mockObj.MockAuthService.EXPECT().CreateToken(gomock.Any(), &entity.UserORM{ID: 1}).
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
