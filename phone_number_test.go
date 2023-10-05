package gosms

import (
	"reflect"
	"testing"
)

func TestPhoneNumber_GetCode(t *testing.T) {
	type fields struct {
		Number int
		Code   int
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
			p := &PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.GetCode(); got != tt.want {
				t.Errorf("GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_GetNumber(t *testing.T) {
	type fields struct {
		Number int
		Code   int
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
			p := &PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.GetNumber(); got != tt.want {
				t.Errorf("GetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_GetUniversalNumber(t *testing.T) {
	type fields struct {
		Number int
		Code   int
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
			p := &PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.GetUniversalNumber(); got != tt.want {
				t.Errorf("GetUniversalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_GetZeroPrefixedNumber(t *testing.T) {
	type fields struct {
		Number int
		Code   int
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
			p := &PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.GetZeroPrefixedNumber(); got != tt.want {
				t.Errorf("GetZeroPrefixedNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_I(t *testing.T) {
	type fields struct {
		Number int
		Code   int
	}
	tests := []struct {
		name   string
		fields fields
		want   IPhoneNumber
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.I(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_getPrefixedIDDCode(t *testing.T) {
	type fields struct {
		Number int
		Code   int
	}
	type args struct {
		prefix string
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
			p := &PhoneNumber{
				Number: tt.fields.Number,
				Code:   tt.fields.Code,
			}
			if got := p.getPrefixedIDDCode(tt.args.prefix); got != tt.want {
				t.Errorf("getPrefixedIDDCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
