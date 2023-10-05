package gosms

import (
	"reflect"
	"testing"
)

func TestMessageContent(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want IMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MessageContent(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MessageContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageTemplate(t *testing.T) {
	type args struct {
		template string
		data     MapStrings
	}
	tests := []struct {
		name string
		args args
		want IMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MessageTemplate(tt.args.template, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MessageTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetContent(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	type args struct {
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetContent(tt.args.gateway); got != tt.want {
				t.Errorf("GetContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetData(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	type args struct {
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MapStrings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetData(tt.args.gateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetGateways(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	tests := []struct {
		name   string
		fields fields
		want   []IGateway
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetGateways(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGateways() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetMessageType(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetMessageType(); got != tt.want {
				t.Errorf("GetMessageType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetSignName(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	type args struct {
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetSignName(tt.args.gateway); got != tt.want {
				t.Errorf("GetSignName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetTemplate(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	type args struct {
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.GetTemplate(tt.args.gateway); got != tt.want {
				t.Errorf("GetTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_I(t *testing.T) {
	type fields struct {
		Type     int
		SignName any
		Content  any
		Template any
		Data     any
		Gateways []IGateway
	}
	tests := []struct {
		name   string
		fields fields
		want   IMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Message{
				Type:     tt.fields.Type,
				SignName: tt.fields.SignName,
				Content:  tt.fields.Content,
				Template: tt.fields.Template,
				Data:     tt.fields.Data,
				Gateways: tt.fields.Gateways,
			}
			if got := m.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}
