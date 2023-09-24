package imobile

import (
	"context"
	"errors"
	"fmt"
	"github.com/pkg6/go-sms"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//https://www.lmobile.cn/ApiPages/index.html

type Imobile struct {
	AccountId string
	Password  string
	ProductId string
}

func GateWay(account, password, productId string) gosms.IGateway {
	return Imobile{
		AccountId: account,
		Password:  password,
		ProductId: productId,
	}.I()
}

func (g Imobile) I() gosms.IGateway {
	return &g
}

func (g *Imobile) AsName() string {
	return "lmobile"
}

func (g *Imobile) url() *url.URL {
	parse, _ := url.Parse("https://api.51welink.com/EncryptionSubmit/SendSms.ashx")
	return parse
}

func (g *Imobile) buildAccessKey(phoneNos, random, timestamp string) string {
	password := strings.ToUpper(gosms.Md5String(g.Password + "SMmsEncryptKey"))
	accessKey := fmt.Sprintf("AccountId=%s&PhoneNos=%s&Password=%s&Random=%s&Timestamp=%s",
		g.AccountId,
		phoneNos,
		password,
		random,
		timestamp,
	)
	return gosms.Sha256String(accessKey)
}

func (g Imobile) Send(to gosms.IPhoneNumber, message gosms.IMessage) (gosms.SMSResult, error) {
	var resp Response
	req := request{}
	timeNow := time.Now()
	req.AccountId = g.AccountId
	req.Timestamp = int(timeNow.Unix())
	req.Random = req.Timestamp
	req.ProductId = g.ProductId
	req.PhoneNos = strconv.Itoa(to.GetNumber())
	req.Content = message.GetContent(g.I())
	req.AccessKey = g.buildAccessKey(req.PhoneNos, strconv.Itoa(req.Random), strconv.Itoa(req.Timestamp))
	response, err := gosms.Client.PostJson(context.Background(), g.url().String(), req)
	_ = response.Unmarshal(&resp)
	result := gosms.BuildSMSResult(to, message, g.I(), resp)
	if err != nil {
		return result, err
	}
	if resp.Result != "succ" && resp.Reason != "提交成功" {
		return result, errors.New(resp.Reason)
	}
	return result, err
}
