package main

import (
	"OTPVerificationGO/app"
	"fmt"

	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/SendOTP", app.RequestOTP)
	router.HandleFunc("/VerifyOTP", app.ConfirmOTP)

	fmt.Printf("Listening.....")
	http.ListenAndServe(":3000", router)
}
