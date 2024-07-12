package api

import (
	"reflect"
	"testing"

	"github.com/voltgizerz/POS-restaurant/internal/adapters/api/middleware"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
)

func TestNewServer(t *testing.T) {
	type args struct {
		interactor    interactor.APInteractor
		jwtMiddleware middleware.JWTAuth
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "SUCCESS",
			args: args{},
			want: &Server{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.interactor, tt.args.jwtMiddleware); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
