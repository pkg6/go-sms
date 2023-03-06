package smsbao

import (
	"errors"
	gosms "github.com/pkg6/go-sms"
	"net/url"
	"strconv"
)

// https://www.smsbao.com/openapi

type SmsBao struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func GateWay(user, password string) gosms.IGateway {
	return SmsBao{User: user, Password: password}.Clone()
}
func (g SmsBao) Clone() gosms.IGateway {
	return &g
}
func (g *SmsBao) GetName() string {
	return "smsbao"
}
func (g *SmsBao) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	var resp Response
	var m string
	c := message.GetContent(g.Clone())
	hostParse, _ := url.Parse("http://api.smsbao.com")
	if to.GetCode() == 0 || to.GetCode() == 86 {
		m = strconv.Itoa(to.GetNumber())
		hostParse.Path = "sms"
	} else {
		m = to.GetUniversalNumber()
		hostParse.Path = "wsms"
	}
	query := gosms.MapStrings{
		"u": g.User,
		"p": gosms.Md5String(g.Password),
		"m": m,
		"c": c,
	}
	response, err := gosms.NewClient(hostParse.String()).SetQueryParams(query).Get(nil)
	result := gosms.BuildSMSResult(to, message, g.Clone(), resp)
	resp.Code = string(response)
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
