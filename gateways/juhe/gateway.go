package juhe

import (
	"errors"
	"github.com/pkg6/go-requests"
	"github.com/pkg6/go-sms"
	"net/url"
)

//https://www.juhe.cn/docs/api/id/54

type JuHe struct {
	Key   string `json:"key"`
	DType string `json:"dtype"`
	gosms.Lock
}

func GateWay(key string) gosms.IGateway {
	return JuHe{Key: key}.I()
}

func (g JuHe) I() gosms.IGateway {
	if g.DType == "" {
		g.DType = "json"
	}
	g.LockInit()
	return &g
}

func (g JuHe) AsName() string {
	return "juhe"
}

func (g JuHe) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	g.Lock.L.Lock()
	defer g.L.Unlock()
	var resp = Response{}
	template := message.GetTemplate(g.I())
	data := message.GetData(g.I())
	param := url.Values{}
	// 接收短信的手机号码
	param.Set("mobile", gosms.GetPhoneNumber(to))
	// 短信模板ID，请参考个人中心短信模板设置
	param.Set("tpl_id", template)
	// 模板变量，如无则不用填写
	vars, err := data.ToJson()
	param.Set("vars", vars)
	// 接口请求Key
	param.Set("key", g.Key)
	response, err := requests.PostForm("http://v.juhe.cn/sms/send", param)
	defer response.Close()
	err = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g, resp)
	if err != nil {
		return result, err
	}
	if resp.ErrorCode != 0 {
		return result, errors.New(resp.Reason)
	}
	return result, err
}
