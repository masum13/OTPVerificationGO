package app

import (
	"errors"
	"strings"

	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

func sendOTP(otpData OTPData) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(otpData.PhoneNumber)
	params.SetChannel("sms")

	client, serviceID, err := TwilioClient()
	if err != nil {
		return "", err
	}

	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {
		if strings.Contains(err.Error(), "ApiError 60200") {
			err = errors.New("please enter country code in the phone number")
		}
		return "", err
	}

	return *resp.Status, nil
}

func VerifyOtp(otpNew VerifyOTP) (string, error) {
	client, serviceID, err := TwilioClient()
	if err != nil {
		return "", err
	}

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(otpNew.PhoneNumber)
	params.SetCode(otpNew.VerificationCode)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)
	if err != nil {
		if strings.Contains(err.Error(), "ApiError 60200") {
			err = errors.New("please add country code in the phone number")
		}
		return "", err
	}

	if *resp.Status != "approved" {
		err := errors.New(*resp.Status)
		return "", err
	}

	return *resp.Status, nil
}
