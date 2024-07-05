package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func RequestOTP(w http.ResponseWriter, r *http.Request) {
	var phone OTPData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&phone)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := ValidateData(&phone); err != nil {
		return
	}

	resp, err := sendOTP(phone)
	if err != nil {
		log.Fatalf("Getting a error to send a OTP: %v", err)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"result": "success", "responce": resp}
	json.NewEncoder(w).Encode(response)
}

func ConfirmOTP(w http.ResponseWriter, r *http.Request) {
	var data VerifyOTP

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := ValidateData(data); err != nil {
		//handle error
		return
	}

	resp, err := VerifyOtp(data)
	if err != nil {
		//handle error
		log.Fatalf("OTP verification failed: %v", err)
		return
	}
	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"result": "success", "responce": resp}
	json.NewEncoder(w).Encode(response)
}
