package api

import (
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestServer_initAuthRoutes(t *testing.T) {
	type args struct {
		group fiber.Router
	}
	tests := []struct {
		name string
		s    *Server
		args args
	}{
		{
			name: "SUCCESS - initAuthRoutes",
			s:    &Server{},
			args: args{group: fiber.New().Group("api/v1")},
		},
	}
	for _, tt := range tests {
		ctrl, mockObj := NewMock(t)
		defer ctrl.Finish()

		tt.s = &Server{
			authHandler: mockObj.MockAuthHandler,
		}

		t.Run(tt.name, func(t *testing.T) {
			tt.s.initAuthRoutes(tt.args.group)
		})
	}
}
