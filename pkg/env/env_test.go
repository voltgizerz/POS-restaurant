package env

import (
	"os"
	"testing"
)

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		name          string
		setEnv        bool   // whether to set the environment variable
		envValue      string // value to set for the environment variable
		expectedValue bool   // expected result from IsDevelopment
	}{
		{
			name:          "Development mode",
			setEnv:        true,
			envValue:      "development",
			expectedValue: true,
		},
		{
			name:          "Non-development mode",
			setEnv:        true,
			envValue:      "production",
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set or unset the environment variable based on the test case
			if tt.setEnv {
				os.Setenv("GO_ENV", tt.envValue)
			} else {
				os.Unsetenv("GO_ENV")
			}

			// Run the IsDevelopment function and compare the result with the expected value
			if got := IsDevelopment(); got != tt.expectedValue {
				t.Errorf("IsDevelopment() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name          string
		setEnv        bool   // whether to set the environment variable
		envValue      string // value to set for the environment variable
		expectedValue bool   // expected result from IsProduction
	}{
		{
			name:          "Production mode",
			setEnv:        true,
			envValue:      "production",
			expectedValue: true,
		},
		{
			name:          "Non-production mode",
			setEnv:        true,
			envValue:      "development",
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set or unset the environment variable based on the test case
			if tt.setEnv {
				os.Setenv("GO_ENV", tt.envValue)
			} else {
				os.Unsetenv("GO_ENV")
			}

			// Run the IsProduction function and compare the result with the expected value
			if got := IsProduction(); got != tt.expectedValue {
				t.Errorf("IsProduction() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}
