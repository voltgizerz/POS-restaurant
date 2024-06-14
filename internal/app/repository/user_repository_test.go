package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

func TestUserRepository_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		r       *UserRepository
		args    args
		want    *entity.UserORM
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
