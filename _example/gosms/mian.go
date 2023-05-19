package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways"
	"github.com/pkg6/go-sms/gateways/ihuyi"
)

func ihuyisend() {
	sms := gosms.NewGateways(ihuyi.GateWay("APIID", "APIKEY"))
	//sendAsNames := []string{"ihuyi"}
	_, e := sms.Send(18888888888, gosms.MapStringAny{
		"content": "您的验证码是：****。请不要把验证码泄露给其他人。",
		"data": gosms.MapStrings{
			"code": "12345",
		},
	}, nil)
	//_, e := g.Send(18888888888, gosms.MapStringAny{
	//	"content": "您的验证码是：****。请不要把验证码泄露给其他人。",
	//	"data": map[string]string{
	//		"code": "12345",
	//	},
	//}, nil)
	fmt.Println(e)
}

func ihuyisend2() {
	sms := gosms.NewParser(gateways.Gateways{IHuYi: ihuyi.IHuYi{Account: "APIID", Password: "APIKEY"}})
	//sendAsNames := []string{"ihuyi"}
	r, e := sms.Send(18888888888, gosms.MapStringAny{
		"content": "您的验证码是：****。请不要把验证码泄露给其他人。",
	}, nil)
	fmt.Println(r, e)
}

func ihuyisendfun() {
	sms := gosms.NewParser(gateways.Gateways{IHuYi: ihuyi.IHuYi{Account: "APIID", Password: "APIKEY"}})
	//sendAsNames := []string{"ihuyi"}
	//_, _ = g.Send(18888888888, gosms.MapStringAny{
	//	"content": func(gateway gosms.IGateway) string {
	//		return "您的验证码是：****。请不要把验证码泄露给其他人。"
	//	},
	//	"template": func(gateway gosms.IGateway) string {
	//		return "SMS_001"
	//	},
	//	//"data": func(gateway gosms.IGateway) map[string]string {
	//	//	return map[string]string{
	//	//		"code": "1234",
	//	//	}
	//	//},
	//	"data": func(gateway gosms.IGateway) gosms.MapStrings {
	//		return map[string]string{
	//			"code": "1234",
	//		}
	//	},
	//}, nil)
	_, _ = sms.Send(18888888888, gosms.MapStringAny{
		"content": func(gateway gosms.IGateway) string {
			return "您的验证码是：****。请不要把验证码泄露给其他人。"
		},
		"template": func(gateway gosms.IGateway) string {
			if gateway.AsName() == "aliyun" {
				return "TP2818"
			}
			return "SMS_001"
		},
		//"data": func(gateway gosms.IGateway) map[string]string {
		//	return map[string]string{
		//		"code": "1234",
		//	}
		//},
		"data": func(gateway gosms.IGateway) gosms.MapStrings {
			return map[string]string{
				"code": "1234",
			}
		},
	}, nil)
	//fmt.Println(r, e)
}

func main() {
	ihuyisend()
}
