package config

import (
	"os"
	"testing"
)

func Test_getConfigPATH(t *testing.T) {
	tests := []struct {
		name         string
		isProduction bool
		want         string
	}{
		{
			name:         "Development mode",
			isProduction: false,
			want:         ConfigPathDevelopment,
		},
		{
			name:         "Production mode",
			isProduction: true,
			want:         ConfigPathPorduction,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isProduction {
				os.Setenv("GO_ENV", "production")
			} else {
				os.Setenv("GO_ENV", "dev")
			}

			if got := getConfigPATH(); got != tt.want {
				t.Errorf("getConfigPATH() = %v, want %v", got, tt.want)
			}
		})
	}
}
