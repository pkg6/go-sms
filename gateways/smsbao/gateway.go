package smsbao

import (
	"context"
	"errors"
	"github.com/pkg6/go-sms"
	"net/url"
	"strconv"
)

// https://www.smsbao.com/openapi

type SmsBao struct {
	User     string `json:"user"`
	Password string `json:"password"`
	gosms.Lock
}

func GateWay(user, password string) gosms.IGateway {
	return SmsBao{User: user, Password: password}.I()
}
func (g SmsBao) I() gosms.IGateway {
	g.LockInit()
	return &g
}
func (g *SmsBao) AsName() string {
	return "smsbao"
}
func (g *SmsBao) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	var resp Response
	var m string
	c := message.GetContent(g.I())
	host := "http://api.smsbao.com"
	if to.GetCode() == 0 || to.GetCode() == 86 {
		m = strconv.Itoa(to.GetNumber())
		host += "/sms"
	} else {
		m = to.GetUniversalNumber()
		host += "/wsms"
	}
	q := url.Values{}
	q.Set("u", g.User)
	q.Set("p", gosms.Md5String(g.Password))
	q.Set("m", m)
	q.Set("c", c)
	response, err := gosms.Client.Get(context.Background(), host, q)
	result := gosms.BuildSMSResult(to, message, g.I(), resp)
	resp.Code = response.ReadAllString()
	defer response.Close()
	if msg, ok := ErrorStatuses[resp.Code]; ok {
		resp.Message = msg
	} else {
		resp.Message = "未知错误信息"
	}
	if resp.Code != "0" {
		return result, errors.New(resp.Message)
	}
	return result, err
}
