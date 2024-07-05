package app

import (
	"os"

	"github.com/twilio/twilio-go"
)

func TwilioClient() (*twilio.RestClient, string, error) {
	serviceSID, ok := os.LookupEnv("TWILIO_SERVICE_SID")
	if !ok {
		panic("TWILIO_SERVICE_SID could not be found")
	}
	authToken, ok := os.LookupEnv("TWILIO_AUTH_TOKEN")
	if !ok {
		panic("TWILIO_AUTH_TOKEN could not be found")
	}
	accountSID, ok := os.LookupEnv("TWILIO_ACCOUNT_SID")
	if !ok {
		panic("TWILIO_ACCOUNT_SID could not be found")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})
	return client, serviceSID, nil
}
