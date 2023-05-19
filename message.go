package gosms

// MessageSignNameFun 生成签名闭包
type MessageSignNameFun func(gateway IGateway) string

// MessageContentFun 生成消息内容闭包
type MessageContentFun func(gateway IGateway) string

// MessageTemplateFun 生成消息模版闭包
type MessageTemplateFun func(gateway IGateway) string

// MessageDataFun 生成消息模版数据闭包
type MessageDataFun func(gateway IGateway) MapStrings

type Message struct {
	//消息类型
	Type int `json:"type"`
	//签名
	SignName any `json:"sign_name"`
	//消息内容
	Content any `json:"content"`
	//消息模版
	Template any `json:"template"`
	//消息模版数据
	Data any `json:"data"`
	//网关
	Gateways []IGateway `json:"gateways"`
}

func MessageContent(content string) IMessage {
	return &Message{Content: content}
}
func MessageTemplate(template string, data MapStrings) IMessage {
	return &Message{Template: template, Data: data}
}
func (m Message) I() IMessage {
	return &m
}

func (m *Message) GetMessageType() int {
	return m.Type
}

func (m *Message) GetSignName(gateway IGateway) string {
	if signName, ok := m.SignName.(string); ok {
		return signName
	}
	if fn, ok := m.SignName.(func(IGateway) string); ok {
		return fn(gateway)
	}
	panic("SignName must be string OR func(gateway IGateway) string")
}

func (m *Message) GetContent(gateway IGateway) string {
	if content, ok := m.Content.(string); ok {
		return content
	}
	if fn, ok := m.Content.(func(IGateway) string); ok {
		return fn(gateway)
	}
	panic("Content must be string OR func(gateway IGateway) string")
}

func (m *Message) GetTemplate(gateway IGateway) string {
	if template, ok := m.Template.(string); ok {
		return template
	}
	if fn, ok := m.Template.(func(gateway IGateway) string); ok {
		return fn(gateway)
	}
	panic("Template must be string OR func(gateway IGateway) string")
}

func (m *Message) GetData(gateway IGateway) MapStrings {
	if data, ok := m.Data.(MapStrings); ok {
		return data
	}
	if fn, ok := m.Data.(func(IGateway) MapStrings); ok {
		return fn(gateway)
	}
	if fn, ok := m.Data.(func(IGateway) map[string]string); ok {
		return fn(gateway)
	}
	panic("Data must be MapStrings OR func(gateway IGateway) MapStrings")
}

func (m *Message) GetGateways() []IGateway {
	return m.Gateways
}
