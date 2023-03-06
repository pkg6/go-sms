package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"net/url"
)

//https://www.twilio.com/docs/sms/send-messages

type Twilio struct {
	AccountSID        string `json:"account_sid"`
	AuthToken         string `json:"auth_token"`
	TwilioPhoneNumber string `json:"twilio_phone_number"`
}

func GateWay(accountSID, authToken, twilioPhoneNumber string) gosms.IGateway {
	return Twilio{AccountSID: accountSID, AuthToken: authToken, TwilioPhoneNumber: twilioPhoneNumber}.Clone()
}

func (g Twilio) Clone() gosms.IGateway {
	return &g
}

func (g *Twilio) GetName() string {
	return "twilio"
}

func (g *Twilio) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	var resp Response
	var data = url.Values{}
	data.Set("Body", message.GetContent(g.Clone()))
	data.Set("From", g.TwilioPhoneNumber)
	data.Set("To", to.GetUniversalNumber())
	response, err := gosms.NewClient(fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", g.AccountSID)).
		SetContentType(gosms.FormASCIIContentType).
		SetBasicAuth(g.AccountSID, g.AuthToken).PostForm(data)
	_ = json.Unmarshal(response, &resp)
	result := gosms.BuildSMSResult(to, message, g.Clone(), resp)
	if resp.Code != 0 {
		return result, errors.New(resp.Message)
	}
	return result, err
}
