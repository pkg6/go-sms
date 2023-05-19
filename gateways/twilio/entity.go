package twilio

// {"code":20003,"message":"Authentication Error - invalid username","more_info":"https://www.twilio.com/docs/errors/20003","status":401}
// {"body": "Sent from your Twilio trial account - Hello from Twilio", "num_segments": "1", "direction": "outbound-api", "from": "+15673717033", "date_updated": "Sun, 05 Mar 2023 01:58:10 +0000", "price": null, "error_message": null, "uri": "/2010-04-01/Accounts/ACde83114a82076f604e676a7daed9f8a0/Messages/SM221dd62603109ed58ddf6af99600d02f.json", "account_sid": "ACde83114a82076f604e676a7daed9f8a0", "num_media": "0", "to": "+8617752535808", "date_created": "Sun, 05 Mar 2023 01:58:10 +0000", "status": "queued", "sid": "SM221dd62603109ed58ddf6af99600d02f", "date_sent": null, "messaging_service_sid": null, "error_code": null, "price_unit": "USD", "api_version": "2010-04-01", "subresource_uris": {"media": "/2010-04-01/Accounts/ACde83114a82076f604e676a7daed9f8a0/Messages/SM221dd62603109ed58ddf6af99600d02f/Media.json"}}

type Response struct {
	Code                int         `json:"code"`
	Message             string      `json:"message"`
	MoreInfo            string      `json:"more_info"`
	Body                string      `json:"body"`
	NumSegments         string      `json:"num_segments"`
	Direction           string      `json:"direction"`
	From                string      `json:"from"`
	DateUpdated         string      `json:"date_updated"`
	Price               interface{} `json:"price"`
	ErrorMessage        interface{} `json:"error_message"`
	URI                 string      `json:"uri"`
	AccountSid          string      `json:"account_sid"`
	NumMedia            string      `json:"num_media"`
	To                  string      `json:"to"`
	DateCreated         string      `json:"date_created"`
	Status              string      `json:"status"`
	Sid                 string      `json:"sid"`
	DateSent            interface{} `json:"date_sent"`
	MessagingServiceSid interface{} `json:"messaging_service_sid"`
	ErrorCode           interface{} `json:"error_code"`
	PriceUnit           string      `json:"price_unit"`
	APIVersion          string      `json:"api_version"`
	SubresourceUris     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
}
