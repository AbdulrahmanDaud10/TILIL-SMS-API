package tilil

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendBulkSMSHandler(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body
	var requestBody struct {
		Text         string   `json:"text"`
		PhoneNumbers []string `json:"phoneNumbers"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Send bulk SMS and check for errors
	err = SendBulkSMS(requestBody.Text, requestBody.PhoneNumbers)
	if err != nil {
		// Log the error for further analysis
		log.Println("Failed to send SMS:", err)

		// Return a more informative error response
		http.Error(w, "Failed to send SMS", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SMS sent successfully"))
}

func SendSingleSMSHandler(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body
	var requestBody struct {
		Text        string `json:"text"`
		PhoneNumber string `json:"phoneNumber"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Send single SMS
	err = SendSingleSMS(requestBody.Text, requestBody.PhoneNumber)
	if err != nil {
		// Handle error gracefully, potentially logging it
		http.Error(w, "Failed to send SMS", http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SMS sent successfully"))
}
