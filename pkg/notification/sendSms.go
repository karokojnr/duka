package notification

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/karokojnr/duka/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSms(phone string, msg string) error
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSms(phone string, msg string) error {
	accountSid := c.config.TwilioAccountSid
	authToken := c.config.TwilioAuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(c.config.TwilioPhoneNumber)
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return errors.New("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}
