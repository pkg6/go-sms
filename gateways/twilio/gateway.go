package twilio

import (
	"context"
	"errors"
	"fmt"
	"github.com/pkg6/go-sms"
	"net/url"
)

//https://www.twilio.com/docs/sms/send-messages

type Twilio struct {
	AccountSID        string `json:"account_sid"`
	AuthToken         string `json:"auth_token"`
	TwilioPhoneNumber string `json:"twilio_phone_number"`
	gosms.Lock
}

func GateWay(accountSID, authToken, twilioPhoneNumber string) gosms.IGateway {
	return Twilio{AccountSID: accountSID, AuthToken: authToken, TwilioPhoneNumber: twilioPhoneNumber}.I()
}

func (g Twilio) I() gosms.IGateway {
	g.LockInit()
	return &g
}

func (g *Twilio) AsName() string {
	return "twilio"
}

func (g *Twilio) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	var resp Response
	var data = url.Values{}
	data.Set("Body", message.GetContent(g.I()))
	data.Set("From", g.TwilioPhoneNumber)
	data.Set("To", to.GetUniversalNumber())
	uri := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", g.AccountSID)
	client := gosms.Client
	client.WithBasicAuth(g.AccountSID, g.AuthToken)
	response, err := client.PostForm(context.Background(), uri, data)
	defer response.Close()
	err = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g, resp)
	if err != nil {
		return result, err
	}
	if resp.Code != 0 {
		return result, errors.New(resp.Message)
	}
	return result, err
}
