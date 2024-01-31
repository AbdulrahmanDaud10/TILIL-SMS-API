package tilil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	Endpoint = os.Getenv("TILIL_SMS_ENDPOINT")
	ApiKey   = os.Getenv("TILIL_API_KEY")
	senderID = os.Getenv("SENDER_ID")
)

type SMSRequestBody struct {
	API_Key   string `json:"api_key"`
	ShortCode string `json:"short_code"`
	Message   string `json:"message"`
	Mobile    string `json:"mobile"`
	ServiceID int    `json:"service_id"`
}

func SendSingleSMS(text, phoneNumber string) error {
	body := SMSRequestBody{
		API_Key:   ApiKey,
		ShortCode: senderID,
		Message:   "How are you?",
		Mobile:    "+254795927076",
		ServiceID: 0,
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	responce, err := http.Post(Endpoint, "aplication/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	responceBody, err := io.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responceBody))
	return nil
}

func SendBulkSMS(text string, phoneNumbers []string) error{
	bulkMessages := []map[string]string{}
	for _, phoneNumber := range phoneNumbers {
		bulkMessages = append(bulkMessages, map[string]string{
			phoneNumber: "mobile",
			text:        "message",
		})
	}

	SMSRequestBody := map[string]interface{}{
		"api_key":   ApiKey,
		"serviceId": 0,
		"shortcode": senderID,
		"messages":  bulkMessages,
	}

	smsBody, err := json.Marshal(SMSRequestBody)
	if err != nil {
		panic(err)
	}

	responce, err := http.Post(Endpoint, "aplication/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	responceBody, err := io.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responceBody))
	return nil
}