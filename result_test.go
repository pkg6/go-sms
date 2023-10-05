package gosms

import (
	"reflect"
	"testing"
)

func TestBuildSMSResult(t *testing.T) {
	type args struct {
		phoneNumber IPhoneNumber
		message     IMessage
		gateway     IGateway
		response    any
	}
	tests := []struct {
		name string
		args args
		want SMSResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildSMSResult(tt.args.phoneNumber, tt.args.message, tt.args.gateway, tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildSMSResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
