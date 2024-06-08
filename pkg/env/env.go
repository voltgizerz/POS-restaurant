package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

const (
	GO_ENV_KEY = "GO_ENV"
)

// LoadENV - load env file.
func LoadENV() {
	if err := godotenv.Load(); err != nil {
		logger.LogStdOut.Warn("No .env file found")
	}
}

func IsDevelopment() bool {
	val, ok := os.LookupEnv(GO_ENV_KEY)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GO_ENV_KEY)
	}

	return val == "development"
}

func IsProduction() bool {
	val, ok := os.LookupEnv(GO_ENV_KEY)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GO_ENV_KEY)
	}

	return val == "production"
}

func GetENV() string {
	val, ok := os.LookupEnv(GO_ENV_KEY)
	if !ok {
		logger.LogStdErr.Fatalf("%s not set\n", GO_ENV_KEY)
	}

	return val
}
