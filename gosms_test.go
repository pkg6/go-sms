package gosms

import (
	"reflect"
	"testing"
)

func TestGoSMSInstance(t *testing.T) {
	type args struct {
		gosms GoSMS
	}
	tests := []struct {
		name string
		args args
		want *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoSMSInstance(tt.args.gosms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoSMSInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_AddSendGateway(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.AddSendGateway(tt.args.gateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSendGateway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_Clone(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_Extend(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		name    string
		gateway IGateway
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.Extend(tt.args.name, tt.args.gateway); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_Send(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		to          any
		message     any
		sendAsNames any
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantSmsResults SMSResults
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			gotSmsResults, err := sms.Send(tt.args.to, tt.args.message, tt.args.sendAsNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSmsResults, tt.wantSmsResults) {
				t.Errorf("Send() gotSmsResults = %v, want %v", gotSmsResults, tt.wantSmsResults)
			}
		})
	}
}

func TestGoSMS_formatGateways(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		sendAsNames any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []IGateway
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.formatGateways(tt.args.sendAsNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatGateways() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_formatMessage(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		message any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.formatMessage(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoSMS_formatPhoneNumber(t *testing.T) {
	type fields struct {
		gateways     GatewayMap
		sendGateways []IGateway
		SendAsNames  []string
	}
	type args struct {
		number any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IPhoneNumber
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sms := &GoSMS{
				gateways:     tt.fields.gateways,
				sendGateways: tt.fields.sendGateways,
				SendAsNames:  tt.fields.SendAsNames,
			}
			if got := sms.formatPhoneNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGateways(t *testing.T) {
	type args struct {
		gateways []IGateway
	}
	tests := []struct {
		name string
		args args
		want *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGateways(tt.args.gateways...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGateways() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewParser(t *testing.T) {
	type args struct {
		parser IParser
	}
	tests := []struct {
		name string
		args args
		want *GoSMS
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.parser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
