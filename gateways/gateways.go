package gateways

import (
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/aliyun"
	"github.com/pkg6/go-sms/gateways/ihuyi"
	"github.com/pkg6/go-sms/gateways/juhe"
	"github.com/pkg6/go-sms/gateways/lmobile"
	"github.com/pkg6/go-sms/gateways/smsbao"
	"github.com/pkg6/go-sms/gateways/twilio"
	"github.com/pkg6/go-sms/gateways/yunxin"
)

type Gateways struct {
	IHuYi   ihuyi.IHuYi     `json:"ihuyi"`
	ALiYun  aliyun.ALiYun   `json:"aliyun"`
	JuHe    juhe.JuHe       `json:"juhe"`
	SmsBao  smsbao.SmsBao   `json:"smsbao"`
	YunXin  yunxin.YunXin   `json:"yunxin"`
	Twilio  twilio.Twilio   `json:"twilio"`
	Imobile lmobile.Imobile `json:"lmobile"`
}

func (g Gateways) FormatGateways() (gateways []gosms.IGateway) {
	if !gosms.IsZero(g.IHuYi) {
		gateways = append(gateways, g.IHuYi.I())
	}
	if !gosms.IsZero(g.ALiYun) {
		gateways = append(gateways, g.ALiYun.I())
	}
	if !gosms.IsZero(g.JuHe) {
		gateways = append(gateways, g.JuHe.I())
	}
	if !gosms.IsZero(g.SmsBao) {
		gateways = append(gateways, g.SmsBao.I())
	}
	if !gosms.IsZero(g.YunXin) {
		gateways = append(gateways, g.YunXin.I())
	}
	if !gosms.IsZero(g.YunXin) {
		gateways = append(gateways, g.YunXin.I())
	}
	if !gosms.IsZero(g.Imobile) {
		gateways = append(gateways, g.Imobile.I())
	}
	return gateways
}
