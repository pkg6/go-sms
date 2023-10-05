package gosms

import (
	"reflect"
	"testing"
)

func TestSender(t *testing.T) {
	type args struct {
		number  IPhoneNumber
		message IMessage
		gateway IGateway
	}
	tests := []struct {
		name    string
		args    args
		want    SMSResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sender(tt.args.number, tt.args.message, tt.args.gateway)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sender() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSenderGateways(t *testing.T) {
	type args struct {
		number   IPhoneNumber
		message  IMessage
		gateways []IGateway
	}
	tests := []struct {
		name          string
		args          args
		wantSmsResult SMSResults
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSmsResult, err := SenderGateways(tt.args.number, tt.args.message, tt.args.gateways)
			if (err != nil) != tt.wantErr {
				t.Errorf("SenderGateways() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSmsResult, tt.wantSmsResult) {
				t.Errorf("SenderGateways() gotSmsResult = %v, want %v", gotSmsResult, tt.wantSmsResult)
			}
		})
	}
}

func TestSenderNumbers(t *testing.T) {
	type args struct {
		numbers []IPhoneNumber
		message IMessage
		gateway IGateway
	}
	tests := []struct {
		name          string
		args          args
		wantSmsResult SMSResults
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSmsResult, err := SenderNumbers(tt.args.numbers, tt.args.message, tt.args.gateway)
			if (err != nil) != tt.wantErr {
				t.Errorf("SenderNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSmsResult, tt.wantSmsResult) {
				t.Errorf("SenderNumbers() gotSmsResult = %v, want %v", gotSmsResult, tt.wantSmsResult)
			}
		})
	}
}
