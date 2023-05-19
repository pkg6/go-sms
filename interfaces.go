package gosms

import "reflect"

// IPhoneNumber 手机号
type IPhoneNumber interface {
	I() IPhoneNumber
	// GetCode 手机号号码前缀 例：中国 86
	GetCode() int
	// GetNumber 手机号 18888888888
	GetNumber() int
	// GetUniversalNumber 返回 +8618888888888
	GetUniversalNumber() string
	// GetZeroPrefixedNumber 返回 008618888888888
	GetZeroPrefixedNumber() string
}

// GatewayName 获取Gateway name 如果不为空就是包的名字
func GatewayName(gateway IGateway) string {
	if name := gateway.AsName(); name != "" {
		return name
	}
	return reflect.TypeOf(gateway).Elem().String()
}

type GatewayMap map[string]IGateway

// IGateway 网关
type IGateway interface {
	// I 用于初始化默认值
	I() IGateway
	// AsName 网关名称
	AsName() string
	// Send 发送操作
	Send(to IPhoneNumber, message IMessage) (SMSResult, error)
}

// IMessage 消息内容
type IMessage interface {
	I() IMessage
	// GetMessageType 消息类型
	GetMessageType() int
	// GetSignName 签名
	GetSignName(gateway IGateway) string
	// GetContent 文字内容，使用在像云片类似的以文字内容发送的平台
	GetContent(gateway IGateway) string
	// GetTemplate 模板 ID，使用在以模板ID来发送短信的平台
	GetTemplate(gateway IGateway) string
	// GetData 模板变量，使用在以模板ID来发送短信的平台
	GetData(gateway IGateway) MapStrings
	// GetGateways 消息指定渠道发送
	GetGateways() []IGateway
}

// IParser 配置信息解析器
type IParser interface {
	// FormatGateways 根据配置可以获取到对应到注册网关
	FormatGateways() (gateways []IGateway)
}
