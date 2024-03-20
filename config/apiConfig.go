package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Port      string
	Dsn       string
	AppSecret string
}

func SetUpEnvironment() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			return AppConfig{}, err
		}
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("port not found in environment variables")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("dsn not found in environment variables")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("app secret not found in environment variables")
	}

	return AppConfig{Port: httpPort, Dsn: dsn, AppSecret: appSecret}, nil
}
