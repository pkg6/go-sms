package ihuyi

import (
	gosms "github.com/pkg6/go-sms"
	"net/url"
	"reflect"
	"testing"
)

func TestGateWay(t *testing.T) {
	type args struct {
		account  string
		password string
	}
	tests := []struct {
		name string
		args args
		want gosms.IGateway
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GateWay(tt.args.account, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GateWay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIHuYi_AsName(t *testing.T) {
	type fields struct {
		Account  string
		Password string
		Format   string
		Lock     gosms.Lock
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &IHuYi{
				Account:  tt.fields.Account,
				Password: tt.fields.Password,
				Format:   tt.fields.Format,
				Lock:     tt.fields.Lock,
			}
			if got := g.AsName(); got != tt.want {
				t.Errorf("AsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIHuYi_I(t *testing.T) {
	type fields struct {
		Account  string
		Password string
		Format   string
		Lock     gosms.Lock
	}
	tests := []struct {
		name   string
		fields fields
		want   gosms.IGateway
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := IHuYi{
				Account:  tt.fields.Account,
				Password: tt.fields.Password,
				Format:   tt.fields.Format,
				Lock:     tt.fields.Lock,
			}
			if got := g.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIHuYi_Send(t *testing.T) {
	type fields struct {
		Account  string
		Password string
		Format   string
		Lock     gosms.Lock
	}
	type args struct {
		to      gosms.IPhoneNumber
		message gosms.IMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    gosms.SMSResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &IHuYi{
				Account:  tt.fields.Account,
				Password: tt.fields.Password,
				Format:   tt.fields.Format,
				Lock:     tt.fields.Lock,
			}
			got, err := g.Send(tt.args.to, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIHuYi_url(t *testing.T) {
	type fields struct {
		Account  string
		Password string
		Format   string
		Lock     gosms.Lock
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &IHuYi{
				Account:  tt.fields.Account,
				Password: tt.fields.Password,
				Format:   tt.fields.Format,
				Lock:     tt.fields.Lock,
			}
			if got := g.url(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("url() = %v, want %v", got, tt.want)
			}
		})
	}
}
