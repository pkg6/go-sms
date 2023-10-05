package gosms

import (
	"reflect"
	"testing"
)

func TestCHNPhoneNumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want IPhoneNumber
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CHNPhoneNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CHNPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPhoneNumber(t *testing.T) {
	type args struct {
		number IPhoneNumber
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPhoneNumber(tt.args.number); got != tt.want {
				t.Errorf("GetPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoCodePhoneNumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want IPhoneNumber
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoCodePhoneNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoCodePhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
