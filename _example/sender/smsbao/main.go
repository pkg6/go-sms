package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/smsbao"
)

func main() {
	var number = gosms.NoCodePhoneNumber(18888888888)
	//对接相关资源 APIID 为account APIKEY 为password
	var gateway = smsbao.GateWay("APIID", "APIKEY")
	var message = gosms.MessageContent("您的验证码是：****。请不要把验证码泄露给其他人。")
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(smsbao.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
