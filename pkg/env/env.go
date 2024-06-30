package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

const (
	GoEnvKey       = "GO_ENV"
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

// LoadENV - load env file.
func LoadENV() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("No .env file found")
	}
}

func IsDevelopment() bool {
	val, ok := os.LookupEnv(GoEnvKey)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GoEnvKey)
	}

	return val == EnvDevelopment
}

func IsProduction() bool {
	val, ok := os.LookupEnv(GoEnvKey)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GoEnvKey)
	}

	return val == EnvProduction
}

func GetENV() string {
	val, ok := os.LookupEnv(GoEnvKey)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GoEnvKey)
	}

	return val
}
