package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/pkg/env"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

const (
	ConfigPathDevelopment = "./config/yml/config.development.yml"
	ConfigPathProduction  = "./config/yml/config.production.yml"
)

type (
	MainConfig struct {
		App      `yaml:"app"`
		API      `yaml:"api"`
		Database `yaml:"database"`
	}
)

type (
	App struct {
		Name    string `yaml:"name" env-required:"true"`
		Version string `yaml:"version" env-required:"true"`
		GoENV   string `env:"GO_ENV" env-required:"true"`
		Host    string `yaml:"host" env-required:"true"`
	}

	API struct {
		PORT         string `yaml:"port"`
		JWTSecretKey string `env:"JWT_SECRET_KEY" env-required:"true"`
	}

	Database struct {
		Host         string `env:"DATABASE_HOST" env-required:"true"`
		PORT         string `env:"DATABASE_PORT" env-required:"true"`
		Username     string `env:"DATABASE_USERNAME" env-required:"true"`
		Password     string `env:"DATABASE_PASSWORD" env-required:"true"`
		Name         string `env:"DATABASE_NAME" env-required:"true"`
		MaxOpenConns int    `yaml:"max_open_conns" env-required:"true"`
		MaxIdleConns int    `yaml:"max_idle_conns" env-required:"true"`
	}
)

// NewConfig returns app config.
func NewConfig() *MainConfig {
	env.LoadENV()

	cfg := &MainConfig{}
	err := cleanenv.ReadConfig(getConfigPATH(), cfg)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"error": err,
		}).Fatalf("[NewConfig] error on ReadConfig")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"error": err,
		}).Fatalf("[NewConfig] error on ReadEnv")
	}

	return cfg
}

func getConfigPATH() string {
	if env.IsProduction() {
		return ConfigPathProduction
	}

	return ConfigPathDevelopment
}
