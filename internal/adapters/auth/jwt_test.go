package auth

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
)

func TestAuthJWT_CreateToken(t *testing.T) {
	type args struct {
		ctx  context.Context
		user models.UserORM
	}
	tests := []struct {
		name    string
		a       *AuthJWT
		args    args
		want    *entity.CreateTokenResponse
		wantErr bool
	}{
		{
			name: "SUCCES - CreateToken",
			a: &AuthJWT{
				SecretKey:      "secret-key-mock",
				ExpireDuration: 24,
			},
			args: args{
				ctx: context.Background(),
				user: models.UserORM{
					ID: 1,
				},
			},
			want:    &entity.CreateTokenResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.CreateToken(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthJWT.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Token == "" {
				t.Errorf("AuthJWT.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthJWT_VerifyToken(t *testing.T) {
	type args struct {
		ctx         context.Context
		tokenString string
	}
	tests := []struct {
		name    string
		a       *AuthJWT
		args    args
		want    *jwt.Token
		want1   jwt.MapClaims
		wantErr bool
	}{
		{
			name: "ERROR - VerifyToken",
			a: &AuthJWT{
				SecretKey:      "secret-key-mock",
				ExpireDuration: 24,
			},
			args: args{
				ctx:         context.Background(),
				tokenString: "jwt",
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.a.VerifyToken(tt.args.ctx, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthJWT.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthJWT.VerifyToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AuthJWT.VerifyToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
