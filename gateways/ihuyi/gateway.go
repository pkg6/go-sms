package ihuyi

import (
	"errors"
	"fmt"
	"github.com/pkg6/go-sms"
	"net/url"
)

//doc  https://www.ihuyi.com/api/sms.html

type IHuYi struct {
	Account  string `json:"account" xml:"account"`
	Password string `json:"password" xml:"password"`
	Format   string `json:"format" xml:"format"`
}

func GateWay(account, password string) gosms.IGateway {
	return IHuYi{
		Account:  account,
		Password: password,
	}.Clone()
}
func (g IHuYi) Clone() gosms.IGateway {
	if g.Format == "" {
		g.Format = "json"
	}
	return &g
}

// GetName 网关名称
func (g *IHuYi) GetName() string {
	return "ihuyi"
}

// Send 发送短信方法
func (g *IHuYi) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	resp := Response{}
	time := gosms.TimeString()
	body := url.Values{}
	content := message.GetContent(g.Clone())
	mobile := gosms.GetPhoneNumber(to)
	account := g.Account
	body.Set("account", account)
	body.Set("password", gosms.Md5String(fmt.Sprintf("%v%v%v%v%v", account, g.Password, mobile, content, time)))
	body.Set("mobile", mobile)
	body.Set("content", content)
	body.Set("time", time)
	err := gosms.Client{
		Url:   "http://106.ihuyi.com/webservice/sms.php",
		Query: gosms.MapStrings{"method": "Submit", "format": g.Format},
		Debug: false,
	}.Clone().PostForm(&resp, body, nil)
	result := gosms.BuildSMSResult(to, message, g.Clone(), resp)
	if resp.Code != 2 {
		return result, errors.New(resp.Msg)
	}
	return result, err
}
