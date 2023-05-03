package tilil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TILIL_SMS_ENDPOINT = "https://api.tililtech.com/sms/v3/sendsms"
	TILIL_API_KEY      = "EQxcdrnes0T6ytoXm4i5IvMKwL1S9VbaqjBFpgzJhDuRUlWAZY3N28GP7CHOkf"
	SENDER_ID          = "2513"
)

type SMSRequestBody struct {
	API_Key   string `json:"api_key"`
	ShortCode string `json:"short_code"`
	Message   string `json:"message"`
	Mobile    string `json:"mobile"`
	ServiceID int    `json:"service_id"`
}

func SendSingleSMS(text, phoneNumber string) {
	body := SMSRequestBody{
		API_Key:   TILIL_API_KEY,
		ShortCode: SENDER_ID,
		Message:   "How are you?",
		Mobile:    "+254795927076",
		ServiceID: 0,
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	responce, err := http.Post(TILIL_SMS_ENDPOINT, "aplication/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	responceBody, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responceBody))
}

func SendBulkSMS(text string, phoneNumbers []string) {
	bulkMessages := []map[string]string{}
	for _, phoneNumber := range phoneNumbers {
		bulkMessages = append(bulkMessages, map[string]string{
			phoneNumber: "mobile",
			text:        "message",
		})
	}

	SMSRequestBody := map[string]interface{}{
		"api_key":   TILIL_API_KEY,
		"serviceId": 0,
		"shortcode": SENDER_ID,
		"messages":  bulkMessages,
	}

	smsBody, err := json.Marshal(SMSRequestBody)
	if err != nil {
		panic(err)
	}

	responce, err := http.Post(TILIL_SMS_ENDPOINT, "aplication/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	responceBody, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responceBody))
}
