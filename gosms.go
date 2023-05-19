package gosms

import (
	"errors"
	"log"
	"strconv"
	"sync"
)

var (
	SingletonGoSMS *GoSMS
	once           = &sync.Once{}
)

type GoSMS struct {
	//注册网关
	gateways GatewayMap
	//处理注册网关之后要发送的网关
	sendGateways []IGateway
	//发送网关的别名
	SendAsNames []string
}

// GoSMSInstance 单例模式
func GoSMSInstance(gosms GoSMS) *GoSMS {
	if SingletonGoSMS == nil {
		once.Do(func() {
			SingletonGoSMS = gosms.Clone()
		})
	}
	return SingletonGoSMS
}

// NewParser  携带解析器实例
func NewParser(parser IParser) *GoSMS {
	sms := GoSMS{}.Clone()
	gateways := parser.FormatGateways()
	for _, gateway := range gateways {
		sms.Extend(gateway.AsName(), gateway)
	}
	return sms
}

// NewGateways 携带网关实例
func NewGateways(gateways ...IGateway) *GoSMS {
	sms := GoSMS{}.Clone()
	for _, gateway := range gateways {
		sms.Extend("", gateway)
	}
	return sms
}

func (sms GoSMS) Clone() *GoSMS {
	sms.gateways = GatewayMap{}
	sms.sendGateways = []IGateway{}
	sms.SendAsNames = []string{}
	return &sms
}

// Extend 注册网关
func (sms *GoSMS) Extend(name string, gateway IGateway) *GoSMS {
	var gateways GatewayMap
	if sms.gateways == nil {
		gateways = make(GatewayMap)
	} else {
		gateways = sms.gateways
	}
	if name == "" {
		name = GatewayName(gateway)
	}
	sms.SendAsNames = append(sms.SendAsNames, name)
	gateways[name] = gateway
	sms.gateways = gateways
	return sms
}

// AddSendGateway 直接注册发送网关
func (sms *GoSMS) AddSendGateway(gateway IGateway) *GoSMS {
	sms.sendGateways = append(sms.sendGateways, gateway)
	return sms
}

// Send 批量发送
func (sms *GoSMS) Send(to, message any, sendAsNames any) (smsResults SMSResults, err error) {
	m := sms.formatMessage(message)
	n := sms.formatPhoneNumber(to)
	gs := sms.formatGateways(sendAsNames)
	if len(gs) <= 0 {
		return smsResults, errors.New("SendGateways not find")
	}
	return SenderGateways(n, m, gs)
}

// 解析手机号码
func (sms *GoSMS) formatPhoneNumber(number any) IPhoneNumber {
	if phoneNumber, ok := number.(IPhoneNumber); ok {
		return phoneNumber
	}
	if phoneNumber, ok := number.(int); ok {
		return NoCodePhoneNumber(phoneNumber)
	}
	if phoneNumber, ok := number.(string); ok {
		n, err := strconv.Atoi(phoneNumber)
		if err != nil {
			log.Fatalf("GoSMS.FormatPhoneNumber strconv.Atoi err=%v", err)
		}
		return NoCodePhoneNumber(n)
	}
	panic("GoSMS.FormatPhoneNumber must be IPhoneNumber or int or string")
}

//   解析任意message消息内容
func (sms *GoSMS) formatMessage(message any) IMessage {
	if m, ok := message.(IMessage); ok {
		for _, gateway := range m.GetGateways() {
			sms.Extend(GatewayName(gateway), gateway)
		}
		return m
	}

	if stringAnyMap, ok := message.(MapStringAny); ok {
		iMessage := stringAnyMap.ToIMessage()
		for _, gateway := range iMessage.GetGateways() {
			sms.Extend(GatewayName(gateway), gateway)
		}
		return iMessage.I()
	}
	if msa, ok := message.(map[string]any); ok {
		stringAnyMap := MapStringAny{}
		for k, v := range msa {
			stringAnyMap.Set(k, v)
		}
		iMessage := stringAnyMap.ToIMessage()
		for _, gateway := range iMessage.GetGateways() {
			sms.Extend(GatewayName(gateway), gateway)
		}
		return iMessage.I()
	}
	panic("GoSMS.FormatMessage must be IMessage or MapStringAny")
}

// 解析网关
func (sms *GoSMS) formatGateways(sendAsNames any) []IGateway {
	if sendAsNames == nil {
		sendAsNames = sms.SendAsNames
	}
	if gateway, ok := sendAsNames.(IGateway); ok {
		sms.AddSendGateway(gateway)
	}
	if gateways, ok := sendAsNames.(GatewayMap); ok {
		for _, gateway := range gateways {
			sms.AddSendGateway(gateway)
		}
	}
	if asNames, ok := sendAsNames.([]string); ok {
		for _, name := range asNames {
			if gateway, ok := sms.gateways[name]; ok {
				sms.AddSendGateway(gateway)
			}
		}
	}
	return sms.sendGateways
}
