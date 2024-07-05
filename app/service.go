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
		} else if strings.Contains(err.Error(), "ApiError 20404") {
			err = errors.New("OTP is expired, Generate New one")
		} else if strings.Contains(err.Error(), "ApiError 60202") {
			err = errors.New("max check attempts reached")
		} else if strings.Contains(err.Error(), "pending") {
			err = errors.New("please enter correct OTP")
		}
		return "", err
	}

	if *resp.Status != "approved" {
		err := errors.New(*resp.Status)
		return "", err
	}

	return *resp.Status, nil
}
