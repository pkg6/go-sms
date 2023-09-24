package ihuyi

import (
	"context"
	"errors"
	"fmt"
	"github.com/pkg6/go-requests"
	"github.com/pkg6/go-sms"
	"net/url"
)

//doc  https://www.ihuyi.com/api/sms.html

type IHuYi struct {
	Account  string `json:"account" xml:"account"`
	Password string `json:"password" xml:"password"`
	Format   string `json:"format" xml:"format"`
	gosms.Lock
}

func GateWay(account, password string) gosms.IGateway {
	return IHuYi{
		Account:  account,
		Password: password,
	}.I()
}
func (g IHuYi) I() gosms.IGateway {
	if g.Format == "" {
		g.Format = "json"
	}
	g.LockInit()
	return &g
}

// GetName 网关名称
func (g *IHuYi) AsName() string {
	return "ihuyi"
}
func (g *IHuYi) url() *url.URL {
	v := url.Values{}
	v.Set("method", "Submit")
	v.Set("format", g.Format)
	return requests.Uri("http://106.ihuyi.com/webservice/sms.php", v)
}

// Send 发送短信方法
func (g *IHuYi) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	resp := Response{}
	time := gosms.TimeString()
	body := url.Values{}
	content := message.GetContent(g.I())
	mobile := gosms.GetPhoneNumber(to)
	account := g.Account
	body.Set("account", account)
	body.Set("password", gosms.Md5String(fmt.Sprintf("%v%v%v%v%v", account, g.Password, mobile, content, time)))
	body.Set("mobile", mobile)
	body.Set("content", content)
	body.Set("time", time)
	response, err := gosms.Client.PostForm(context.Background(), g.url().String(), body)
	defer response.Close()
	err = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g, resp)
	if err != nil {
		return result, err
	}
	if resp.Code != 2 {
		return result, errors.New(resp.Msg)
	}
	return result, err
}
