package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Port              string
	Dsn               string
	AppSecret         string
	TwilioAccountSid  string
	TwilioAuthToken   string
	TwilioPhoneNumber string
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

	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("twilio account sid not found in environment variables")
	}

	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("twilio auth token not found in environment variables")
	}

	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("twilio phone number not found in environment variables")
	}

	return AppConfig{
		Port:              httpPort,
		Dsn:               dsn,
		AppSecret:         appSecret,
		TwilioAccountSid:  twilioAccountSid,
		TwilioAuthToken:   twilioAuthToken,
		TwilioPhoneNumber: twilioPhoneNumber,
	}, nil
}
