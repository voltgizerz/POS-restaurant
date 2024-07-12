package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
)

func TestMenuService_GetMenu(t *testing.T) {
	type args struct {
		ctx    context.Context
		idUser int64
	}
	tests := []struct {
		name    string
		s       *MenuService
		args    args
		want    []*entity.MenuResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMenu(tt.args.ctx, tt.args.idUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("MenuService.GetMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MenuService.GetMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMenuService_RegisterMenu(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuData entity.Menu
	}
	tests := []struct {
		name    string
		s       *MenuService
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterMenu(tt.args.ctx, tt.args.menuData)
			if (err != nil) != tt.wantErr {
				t.Errorf("MenuService.RegisterMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MenuService.RegisterMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}
