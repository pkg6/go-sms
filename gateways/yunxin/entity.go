package yunxin

import gosms "github.com/pkg6/go-sms"

var ErrorStatuses = gosms.MapStrings{
	"200": "操作成功",
	"315": "IP限制",
	"403": "非法操作或没有权限",
	"404": "对象不存在",
	"413": "验证失败(短信服务)",
	"414": "参数错误",
	"500": "服务器内部错误",
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Obj     []struct {
		Status     int    `json:"status"`
		Mobile     string `json:"mobile"`
		Updatetime int64  `json:"updatetime"`
	} `json:"obj"`
}
