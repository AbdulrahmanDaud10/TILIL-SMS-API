package main

import (
	"github.com/abdulrahmandaud10/tilil-sms-api/tilil"
)

func main() {
	phoneNumbers := []string{"0712345678", "0712345678", "0712345678", "0712345678"}

	tilil.SendSingleSMS("Welcome to my github simple bulk sms sending API", "0712345678")
	tilil.SendBulkSMS("Welcome to my github simple bulk sms sending API", phoneNumbers)
}
