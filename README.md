<h1 align="center">Go SMS</h1>

<p align="center">:calling: 一款满足你的多种发送需求的短信发送组件</p>


[![Go Report Card](https://goreportcard.com/badge/github.com/pkg6/go-sms)](https://goreportcard.com/report/github.com/pkg6/go-sms)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/pkg6/go-sms?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/pkg6/go-sms/-/badge.svg)](https://sourcegraph.com/github.com/pkg6/go-sms?badge)
[![Release](https://img.shields.io/github/release/pkg6/go-sms.svg?style=flat-square)](https://github.com/pkg6/go-sms/releases)


## 特点

1. 支持目前市面多家服务商
1. 一套写法兼容所有平台
1. 简单配置即可灵活增减服务商
1. 内置多种服务商轮询策略、支持自定义轮询策略
1. 统一的返回值格式，便于日志与监控
1. 更多等你去发现与改进...

## 平台支持

### 目前支持

- [阿里云](https://www.aliyun.com/)
- [互亿无线](http://www.ihuyi.com)
- [短信宝](http://www.smsbao.com/)
- [聚合数据](https://www.juhe.cn)
- [网易云信](https://yunxin.163.com/sms)
- [微网通联](https://www.lmobile.cn/)
- [twilio](https://www.twilio.com/)

## 环境需求

- Golang >= 1.18


## 安装

```shell
$ go get github.com/pkg6/go-sms
```

## 使用

~~~
package main

import (
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways"
	"github.com/pkg6/go-sms/gateways/twilio"
)

func main() {
	sms := gosms.NewParser(gateways.Gateways{Twilio: twilio.Twilio{AccountSID: "ACd********", AuthToken: "***********", TwilioPhoneNumber: "+1********"}})
	// 常规
	sms.Send(18888888888, gosms.MapStringAny{
		"content":  "您的验证码是：****。请不要把验证码泄露给其他人。",
		"template": "SMS_001",
		"data": gosms.MapStrings{
			"code": "6379",
		},
	}, nil)
}
~~~

## 短信内容

由于使用多网关发送，所以一条短信要支持多平台发送，每家的发送方式不一样，但是我们抽象定义了以下公用属性：

- `content` 文字内容，使用在像云片类似的以文字内容发送的平台
- `sign_name` 签名
- `template` 模板 ID，使用在以模板ID来发送短信的平台
- `data`  模板变量，使用在以模板ID来发送短信的平台

所以，在使用过程中你可以根据所要使用的平台定义发送的内容。

~~~
sms.Send(18888888888, gosms.MapStringAny{
    "content":  "您的验证码是：****。请不要把验证码泄露给其他人。",
    "template": "SMS_001",
    "data": gosms.MapStrings{
        "code": "6379",
    },
}, nil)
~~~

你也可以使用闭包来返回对应的值：
~~~
sms.Send(18888888888, gosms.MapStringAny{
    "content": func(gateway gosms.IGateway) string {
        return "您的验证码是：****。请不要把验证码泄露给其他人。"
    },
    "template": func(gateway gosms.IGateway) string {
        if gateway.GetName() == "aliyun" {
            return "TP2818"
        }
        return "SMS_001"
    },
    "data": func(gateway gosms.IGateway) gosms.MapStrings {
        return map[string]string{
            "code": "1234",
        }
    },
}, nil)
~~~

## 发送网关

默认使用初始化中的网关发送，如果某一条短信你想要覆盖默认的设置。在 `send` 方法中使用第三个参数即可：

```
sms := gosms.NewGateways(twilio.GateWay("ACd********", "***********", "+1***********"))
sms.Send(18888888888, gosms.MapStringAny{
    "content": "您的验证码是：****。请不要把验证码泄露给其他人。",
}, nil)
```

## 自定义网关

本拓展已经支持用户自定义网关，你可以很方便的配置即可当成与其它拓展一样的使用：

~~~
sms := gosms.NewGateways()
sms.Extend("aliyun", aliyun.GateWay("accessKeyId", "accessKeySecret"))
sms.Send(18888888888, gosms.MapStringAny{
    "content": "您的验证码是：****。请不要把验证码泄露给其他人。",
}, []string{"aliyun"})
~~~

> 自定义网关实现接口

~~~
// IGateway 网关
type IGateway interface {
	// I 用于初始化默认值
	I() IGateway
	// GetName 网关名称
	GetName() string
	// Send 发送操作
	Send(to IPhoneNumber, message IMessage) (SMSResult, error)
}
~~~

## 通过Sender函数去发送
~~~
package main

import (
	"fmt"
	gosms "github.com/pkg6/go-sms"
	"github.com/pkg6/go-sms/gateways/twilio"
)

func main() {
	gateway := twilio.GateWay("ACd********", "********", "+1111111")
	var message = gosms.MessageContent("Hello from Twilio")
	number := gosms.CHNPhoneNumber(18888888888)
	result, err := gosms.Sender(number, message, gateway)
	if resp, ok := result.ClientResult.Response.(twilio.Response); ok {
		fmt.Println(resp)
	}
	fmt.Println(err)
}
~~~

## 加入我们

如果你认可我们的开源项目，有兴趣为 go-sms的发展做贡献，竭诚欢迎加入我们一起开发完善。无论是[报告错误](https://github.com/pkg6/go-sms/issues)或是
[Pull Request](https://github.com/pkg6/go-sms/pulls) 开发，那怕是修改一个错别字也是对我们莫大的帮助。