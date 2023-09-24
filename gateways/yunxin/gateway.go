package yunxin

import (
	"context"
	"errors"
	"fmt"
	"github.com/pkg6/go-sms"
	"net/url"
	"strconv"
	"strings"
)

type YunXin struct {
	AppKey    string
	AppSecret string
	gosms.Lock
}

func GateWay(appKey, appSecret string) gosms.IGateway {
	return YunXin{AppKey: appKey, AppSecret: appSecret}.I()
}

func (g YunXin) I() gosms.IGateway {
	g.LockInit()
	return &g
}

func (g YunXin) AsName() string {
	return "yunxin"
}

func (g YunXin) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	var resp Response
	data := message.GetData(g.I())
	action := data.GetDefault("action", "sendCode")
	host := fmt.Sprintf("https://api.netease.im/%s/%s.action", "sms", strings.ToLower(action))
	param := url.Values{}
	switch action {
	case "sendCode":
		templateId := message.GetTemplate(g.I())
		param.Set("mobile", to.GetUniversalNumber())
		param.Set("authCode", data.GetDefault("code", ""))
		param.Set("deviceId", data.GetDefault("device_id", ""))
		param.Set("templateid", templateId)
		param.Set("codeLen", strconv.Itoa(len(data.GetDefault("code", ""))))
		param.Set("needUp", data.GetDefault("need_up", "false"))
	case "verifyCode":
		param.Set("mobile", to.GetUniversalNumber())
		param.Set("code", data.GetDefault("code", ""))
	default:
		err := errors.New(fmt.Sprintf("action:%v not supported", action))
		return gosms.BuildSMSResult(to, message, g.I(), resp), err
	}
	headers := gosms.MapStrings{
		"AppKey":  g.AppKey,
		"Nonce":   gosms.Md5String(gosms.Uniqid("gosms")),
		"CurTime": gosms.TimeString(),
	}
	checkSum := gosms.Sha1String(fmt.Sprintf("%v%v%v",
		g.AppSecret,
		headers.GetDefault("Nonce", ""),
		headers.GetDefault("CurTime", "")),
	)
	client := gosms.Client
	client.WithHeader("CheckSum", checkSum)

	response, err := client.PostForm(context.Background(), host, param)
	err = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g.I(), resp)
	if err != nil {
		return result, err
	}
	if msg, ok := ErrorStatuses[strconv.Itoa(resp.Code)]; ok {
		resp.Message = msg
	} else {
		resp.Message = "未知错误信息"
	}
	if resp.Code != 200 {
		return result, errors.New(resp.Message)
	}
	return result, err
}
