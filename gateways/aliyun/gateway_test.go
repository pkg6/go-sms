package aliyun

import (
	gosms "github.com/pkg6/go-sms"
	"reflect"
	"testing"
)

func TestALiYun_AsName(t *testing.T) {
	type fields struct {
		Host            string
		RegionId        string
		AccessKeyId     string
		AccessKeySecret string
		Format          string
		Lock            gosms.Lock
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
			g := &ALiYun{
				Host:            tt.fields.Host,
				RegionId:        tt.fields.RegionId,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Format:          tt.fields.Format,
				Lock:            tt.fields.Lock,
			}
			if got := g.AsName(); got != tt.want {
				t.Errorf("AsName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestALiYun_I(t *testing.T) {
	type fields struct {
		Host            string
		RegionId        string
		AccessKeyId     string
		AccessKeySecret string
		Format          string
		Lock            gosms.Lock
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
			g := ALiYun{
				Host:            tt.fields.Host,
				RegionId:        tt.fields.RegionId,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Format:          tt.fields.Format,
				Lock:            tt.fields.Lock,
			}
			if got := g.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestALiYun_Send(t *testing.T) {
	type fields struct {
		Host            string
		RegionId        string
		AccessKeyId     string
		AccessKeySecret string
		Format          string
		Lock            gosms.Lock
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
			g := &ALiYun{
				Host:            tt.fields.Host,
				RegionId:        tt.fields.RegionId,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Format:          tt.fields.Format,
				Lock:            tt.fields.Lock,
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

func TestALiYun_generateSign(t *testing.T) {
	type fields struct {
		Host            string
		RegionId        string
		AccessKeyId     string
		AccessKeySecret string
		Format          string
		Lock            gosms.Lock
	}
	type args struct {
		query gosms.MapStrings
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
			g := &ALiYun{
				Host:            tt.fields.Host,
				RegionId:        tt.fields.RegionId,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Format:          tt.fields.Format,
				Lock:            tt.fields.Lock,
			}
			if got := g.generateSign(tt.args.query); got != tt.want {
				t.Errorf("generateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestALiYun_query(t *testing.T) {
	type fields struct {
		Host            string
		RegionId        string
		AccessKeyId     string
		AccessKeySecret string
		Format          string
		Lock            gosms.Lock
	}
	tests := []struct {
		name   string
		fields fields
		want   gosms.MapStrings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &ALiYun{
				Host:            tt.fields.Host,
				RegionId:        tt.fields.RegionId,
				AccessKeyId:     tt.fields.AccessKeyId,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Format:          tt.fields.Format,
				Lock:            tt.fields.Lock,
			}
			if got := g.query(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGateWay(t *testing.T) {
	type args struct {
		accessKeyId     string
		accessKeySecret string
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
			if got := GateWay(tt.args.accessKeyId, tt.args.accessKeySecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GateWay() = %v, want %v", got, tt.want)
			}
		})
	}
}
