package auth

import (
	"context"
	"reflect"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
)

func TestGetUserLoginFromCtx(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		want        *Auth
		wantErr     bool
		expectedErr string
	}{
		{
			name: "Valid context",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(
							context.WithValue(context.Background(), constants.CTXKeyUserID, float64(1)),
							constants.CTXKeyUsername, "testuser"),
						constants.CTXKeyRoleID, float64(2)),
					constants.CTXKeyIsActive, true),
				constants.CTXKeyRequestID, "request123")},
			want: &Auth{
				UserID:       1,
				Username:     "testuser",
				RoleID:       2,
				IsUserActive: true,
				RequestID:    "request123",
			},
			wantErr: false,
		},
		{
			name: "Missing userID",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(context.Background(), constants.CTXKeyUsername, "testuser"),
						constants.CTXKeyRoleID, float64(2)),
					constants.CTXKeyIsActive, true),
				constants.CTXKeyRequestID, "request123")},
			want:        nil,
			wantErr:     true,
			expectedErr: "user ID not found in context",
		},
		{
			name: "Missing username",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(context.Background(), constants.CTXKeyUserID, float64(1)),
						constants.CTXKeyRoleID, float64(2)),
					constants.CTXKeyIsActive, true),
				constants.CTXKeyRequestID, "request123")},
			want:        nil,
			wantErr:     true,
			expectedErr: "username not found in context",
		},
		{
			name: "Missing roleID",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(context.Background(), constants.CTXKeyUserID, float64(1)),
						constants.CTXKeyUsername, "testuser"),
					constants.CTXKeyIsActive, true),
				constants.CTXKeyRequestID, "request123")},
			want:        nil,
			wantErr:     true,
			expectedErr: "role ID not found in context",
		},
		{
			name: "Missing isUserActive",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(context.Background(), constants.CTXKeyUserID, float64(1)),
						constants.CTXKeyUsername, "testuser"),
					constants.CTXKeyRoleID, float64(2)),
				constants.CTXKeyRequestID, "request123")},
			want:        nil,
			wantErr:     true,
			expectedErr: "active status not found in context",
		},
		{
			name: "Missing requestID",
			args: args{ctx: context.WithValue(
				context.WithValue(
					context.WithValue(
						context.WithValue(context.Background(), constants.CTXKeyUserID, float64(1)),
						constants.CTXKeyUsername, "testuser"),
					constants.CTXKeyRoleID, float64(2)),
				constants.CTXKeyIsActive, true)},
			want:        nil,
			wantErr:     true,
			expectedErr: "request ID not found in context",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserLoginFromCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserLoginFromCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("GetUserLoginFromCtx() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserLoginFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
