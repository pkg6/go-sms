package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/juhe"
	"github.com/pkg6/go-sms/gateways/yunxin"
)

func main() {
	var number = gosms.NoCodePhoneNumber(18888888888)
	//对接相关资源 APIID 为account APIKEY 为password
	var gateway = yunxin.GateWay("APIID", "222")
	var message = gosms.MessageTemplate("111", gosms.MapStrings{
		//sendCode 或 verifyCode
		"action": "sendCode",
		"code":   "1234",
	})
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(juhe.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
