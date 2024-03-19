package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/twilio"
)

func main() {
	gateway := twilio.GateWay("ACd********", "********", "+1111111")
	var message = gosms.MessageContent("Hello from Twilio")
	number := gosms.CHNPhoneNumber(18888888888)
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(twilio.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
