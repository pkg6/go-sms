package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/ihuyi"
	"github.com/pkg6/go-sms/gateways/imobile"
	"github.com/pkg6/go-sms/gateways/juhe"
	"github.com/pkg6/go-sms/gateways/smsbao"
	"github.com/pkg6/go-sms/gateways/twilio"
	"github.com/pkg6/go-sms/gateways/yunxin"
)

// ihuyi 发送操作
func ihuyiSend() {
	var number = gosms.NoCodePhoneNumber(18888888888)
	//对接相关资源 APIID 为account APIKEY 为password
	var gateway = ihuyi.GateWay("APIID", "APIKEY")
	var message = gosms.MessageContent("您的验证码是：****。请不要把验证码泄露给其他人。")
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(ihuyi.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}

// 短信宝 发送操作
func smsbaosend() {
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

// 聚合 发送操作
func juhesend() {
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

// 云信 发送操作
func yunXinsend() {
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

func twiliosend() {
	gateway := twilio.GateWay("ACd********", "********", "+1111111")
	var message = gosms.MessageContent("Hello from Twilio")
	number := gosms.CHNPhoneNumber(18888888888)
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(twilio.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}

//https://www.lmobile.cn/
func imobileSend() {
	gateway := imobile.GateWay("服务商提供登录账号", "服务商提供登录密码", "产品名称")
	var message = gosms.MessageContent("【某某科技】您好，您的手机验证码是：089548，请在10分钟内使用。若非本人操作，请忽略！")
	number := gosms.CHNPhoneNumber(18888888888)
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(imobile.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
func main() {
	ihuyiSend()
	smsbaosend()
	juhesend()
	imobileSend()
	yunXinsend()
	twiliosend()
}
