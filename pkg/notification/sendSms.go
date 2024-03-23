package notification

import "github.com/karokojnr/duka/config"

type NotificationClient interface {
	SendSms(phone string, msg string) error
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSms(phone string, msg string) error {
	return nil
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}
