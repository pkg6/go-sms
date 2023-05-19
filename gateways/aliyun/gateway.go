package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"github.com/pkg6/go-requests"
	"github.com/pkg6/go-sms"
	"net/url"
	"strings"
	"time"
)

//https://help.aliyun.com/document_detail/419273.html

type ALiYun struct {
	Host            string `json:"host" xml:"host"`
	RegionId        string `json:"region_id" xml:"region_id"`
	AccessKeyId     string `json:"access_key_id" xml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" xml:"access_key_secret"`
	Format          string `json:"format" xml:"format"`
	gosms.Lock
}

func GateWay(accessKeyId, accessKeySecret string) gosms.IGateway {
	gateway := &ALiYun{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	return gateway.I()
}
func (g ALiYun) I() gosms.IGateway {
	if g.Host == "" {
		g.Host = "http://dysmsapi.aliyuncs.com"
	}
	if g.RegionId == "" {
		g.RegionId = "cn-hangzhou"
	}
	if g.Format == "" {
		g.Format = "JSON"
	}
	g.LockInit()
	return &g
}

func (g *ALiYun) AsName() string {
	return "aliyun"
}

//请求参数生成
func (g *ALiYun) query() gosms.MapStrings {
	if g.RegionId == "" {
		g.RegionId = "cn-hangzhou"
	}
	if g.Format == "" {
		g.Format = "JSON"
	}
	maps := gosms.MapStrings{
		"AccessKeyId":      g.AccessKeyId,
		"Action":           "SendSms",
		"Format":           g.Format,
		"RegionId":         g.RegionId,
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   gosms.Uniqid("gosms"),
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Version":          "2017-05-25",
	}
	return maps
}

func (g *ALiYun) generateSign(query gosms.MapStrings) string {
	s := strings.Replace(query.ToUrlQuery(true), "+", "%20", -1)
	stringToSign := "GET&" + url.QueryEscape("/") + "&" + s
	mac := hmac.New(sha1.New, []byte(g.AccessKeySecret+"&"))
	mac.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return url.QueryEscape(signature)
}

func (g *ALiYun) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	var resp Response
	data := message.GetData(g.I())
	signName := data.GetDefault("signName", message.GetSignName(g.I()))
	template := message.GetTemplate(g.I())
	data.Delete("signName")
	query := g.query()
	mobile := gosms.GetPhoneNumber(to)
	query["PhoneNumbers"] = mobile
	query["SignName"] = signName
	query["TemplateCode"] = template
	jsonStr, _ := data.ToJson()
	query["TemplateParam"] = jsonStr
	query["Signature"] = g.generateSign(query)
	response, err := requests.Get(g.Host, query)
	defer response.Close()
	err = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g, resp)
	if err != nil {
		return result, err
	}
	if resp.Code != "ok" {
		return result, errors.New(resp.Message)
	}
	return result, err
}
