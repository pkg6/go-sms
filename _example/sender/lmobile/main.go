package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/lmobile"
)

func main() {
	gateway := lmobile.GateWay("服务商提供登录账号", "服务商提供登录密码", "产品名称")
	var message = gosms.MessageContent("【某某科技】您好，您的手机验证码是：089548，请在10分钟内使用。若非本人操作，请忽略！")
	number := gosms.CHNPhoneNumber(18888888888)
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(lmobile.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
