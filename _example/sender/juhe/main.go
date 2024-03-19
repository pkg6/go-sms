package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/juhe"
)

func main() {
	var number = gosms.NoCodePhoneNumber(18888888888)
	//对接相关资源 APIID 为account APIKEY 为password
	var gateway = juhe.GateWay("APIID")
	var message = gosms.MessageTemplate("111", gosms.MapStrings{
		"code": "1234",
		"name": "聚合数据",
	})
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(juhe.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
