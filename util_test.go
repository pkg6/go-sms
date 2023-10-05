package gosms

import (
	"reflect"
	"sync"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	type args struct {
		text string
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
			if got := Base64Encode(tt.args.text); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.s); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLock_LockInit(t *testing.T) {
	type fields struct {
		L *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lock{
				L: tt.fields.L,
			}
			l.LockInit()
		})
	}
}

func TestMd5String(t *testing.T) {
	type args struct {
		text string
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
			if got := Md5String(tt.args.text); got != tt.want {
				t.Errorf("Md5String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rand(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Rand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandLetterString(t *testing.T) {
	type args struct {
		lenNum int
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
			if got := RandLetterString(tt.args.lenNum); got != tt.want {
				t.Errorf("RandLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandNumString(t *testing.T) {
	type args struct {
		lenNum int
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
			if got := RandNumString(tt.args.lenNum); got != tt.want {
				t.Errorf("RandNumString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandString(t *testing.T) {
	type args struct {
		lenNum int
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
			if got := RandString(tt.args.lenNum); got != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomBytes(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomBytes(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomInt(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RandomInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha1String(t *testing.T) {
	type args struct {
		text string
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
			if got := Sha1String(tt.args.text); got != tt.want {
				t.Errorf("Sha1String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha256String(t *testing.T) {
	type args struct {
		text string
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
			if got := Sha256String(tt.args.text); got != tt.want {
				t.Errorf("Sha256String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSleep(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sleep(tt.args.t)
		})
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(); got != tt.want {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeString(); got != tt.want {
				t.Errorf("TimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLEncode(t *testing.T) {
	type args struct {
		text string
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
			if got := URLEncode(tt.args.text); got != tt.want {
				t.Errorf("URLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqid(t *testing.T) {
	type args struct {
		prefix string
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
			if got := Uniqid(tt.args.prefix); got != tt.want {
				t.Errorf("Uniqid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsleep(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Usleep(tt.args.t)
		})
	}
}
