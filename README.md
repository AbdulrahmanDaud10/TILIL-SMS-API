# TILIL-SMS-API
A quick step by step guide on how to send bulk and premium messages using an application interface (API). It has been made simple to enable a quick integration into Tilil technologies bulk messaging services.

## BULK MESSAGING
```
Required endpoints:
const (
	    TILIL_SINGLE_SMS_ENDPOINT = "https://api.tililtech.com/sms/v3/sendsms"
	    TILIL_API_KEY             = "EQxcdrnes0T6ytoXm4i5IvMKwL1S9VbaqjBFpgzJhDuRUlWAZY3N28GP7CHOkf"
	    SENDER_ID                 = "2513"
      )
```

Edit the variables to your preference

### SMS Request Body:
Below is a sample sendsms JSON data:

```json
{
    "api_key": "{{Test API Key}}",
    "service_id": 0,
    "mobile": "0708400000",
    "response_type": "json",
    "shortcode": "Tilil",
    "message": "This is a message.\n\nRegards\nTilil"
}
```

### SMS Responses Body:
SMS Response(s):
```json
1.Success
    [
        {
            "status_code": "1000",
            "status_desc": "Success",
            "message_id": 288369252,
            "mobile_number": "254708400000",
            "network_id": "1",
            "message_cost": 1,
            "credit_balance": 148
        }
    ]

2.Failed
    [
        {
            "status_code": "1003",
            "status_desc": "Invalid mobile number",
            "message_id": "0",
            "mobile_number": "123",
            "network_id": "",
            "message_cost": "",
            "credit_balance": ""
        }
    ]
```

## REMINDER

The api key provided above is for a demo purpose, visit the official website https://tililtechnologies.co.ke/ to create an account acquire the neccessary shortcodes, senderID important data and information required.