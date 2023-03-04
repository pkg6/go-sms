package twilio

type Response struct {
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
