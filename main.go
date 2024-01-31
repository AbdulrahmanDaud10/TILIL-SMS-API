package main

import (
	"log"
	"net/http"

	"github.com/abdulrahmandaud10/tilil-sms-api/tilil"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	phoneNumbers := []string{"0712345678", "0712345678", "0712345678", "0712345678"}

	tilil.SendSingleSMS("Welcome to my github simple bulk sms sending API", "0712345678")
	tilil.SendBulkSMS("Welcome to my github simple bulk sms sending API", phoneNumbers)

	http.HandleFunc("/sendmessage", tilil.SendSingleSMSHandler)
	http.HandleFunc("/sendbulkmessage", tilil.SendBulkSMSHandler)
	log.Fatal(http.ListenAndServe(":6969", nil))
}
