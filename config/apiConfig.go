package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Port string
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
		return AppConfig{}, errors.New("port not found")
	}
	return AppConfig{Port: httpPort}, nil
}
