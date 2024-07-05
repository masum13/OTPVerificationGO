package app

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type OTPData struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type VerifyOTP struct {
	VerificationCode string `json:"verification_code" validate:"required"`
	OTPData
}

func ValidateData(data any) error {
	var validate = validator.New()
	if err := validate.Struct(data); err != nil {
		log.Fatalf("Request structure is invalid: %v", err)
	}
	return nil
}
