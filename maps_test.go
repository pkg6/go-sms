package gosms

import (
	"reflect"
	"sync"
	"testing"
)

func TestHashMap_Delete(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			g.Delete(tt.args.key)
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			got, got1 := g.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHashMap_Keys(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			if got := g.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Set(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			g.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestHashMap_ToMaps(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			if got := g.ToMaps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Values(t *testing.T) {
	type fields struct {
		HashMap sync.Map
	}
	tests := []struct {
		name   string
		fields fields
		want   []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &HashMap{
				HashMap: tt.fields.HashMap,
			}
			if got := g.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStringAny_Exist(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		maps MapStringAny
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Exist(tt.args.k); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStringAny_Get(t *testing.T) {
	type args struct {
		k  string
		dv any
	}
	tests := []struct {
		name string
		maps MapStringAny
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Get(tt.args.k, tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStringAny_JsonUnmarshal(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		maps    MapStringAny
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.maps.JsonUnmarshal(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JsonUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMapStringAny_Set(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name string
		maps MapStringAny
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.Set(tt.args.k, tt.args.v)
		})
	}
}

func TestMapStringAny_SetForce(t *testing.T) {
	type args struct {
		k     string
		v     any
		force bool
	}
	tests := []struct {
		name string
		maps MapStringAny
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.SetForce(tt.args.k, tt.args.v, tt.args.force)
		})
	}
}

func TestMapStringAny_ToIMessage(t *testing.T) {
	tests := []struct {
		name string
		maps MapStringAny
		want IMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.ToIMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStringAny_ToJson(t *testing.T) {
	tests := []struct {
		name    string
		maps    MapStringAny
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.maps.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStringAny_Verify(t *testing.T) {
	type args struct {
		mustKey []string
	}
	tests := []struct {
		name string
		maps MapStringAny
		args args
		want MapStrings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Verify(tt.args.mustKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Delete(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.Delete(tt.args.k)
		})
	}
}

func TestMapStrings_Exist(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Exist(tt.args.k); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_GetDefault(t *testing.T) {
	type args struct {
		k string
		d string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.GetDefault(tt.args.k, tt.args.d); got != tt.want {
				t.Errorf("GetDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Implode(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Implode(tt.args.sep); got != tt.want {
				t.Errorf("Implode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Keys(t *testing.T) {
	tests := []struct {
		name string
		maps MapStrings
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Set(t *testing.T) {
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.Set(tt.args.k, tt.args.v)
		})
	}
}

func TestMapStrings_SetForce(t *testing.T) {
	type args struct {
		k     string
		v     string
		force bool
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.SetForce(tt.args.k, tt.args.v, tt.args.force)
		})
	}
}

func TestMapStrings_ToJson(t *testing.T) {
	tests := []struct {
		name    string
		maps    MapStrings
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.maps.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_ToUrlQuery(t *testing.T) {
	type args struct {
		kSort bool
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.ToUrlQuery(tt.args.kSort); got != tt.want {
				t.Errorf("ToUrlQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Values(t *testing.T) {
	tests := []struct {
		name string
		maps MapStrings
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapStrings_Verify(t *testing.T) {
	type args struct {
		mustKey []string
	}
	tests := []struct {
		name string
		maps MapStrings
		args args
		want MapStrings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maps.Verify(tt.args.mustKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeMapsString(t *testing.T) {
	type args struct {
		maps []MapStrings
	}
	tests := []struct {
		name string
		args args
		want MapStrings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMapsString(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMapsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceMapStrings(t *testing.T) {
	type args struct {
		s    string
		news map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "手机号",
			args: args{s: "您手机注册的验证码为：【变量1】，如有问题请拨打客服电话：【变量2】", news: map[string]string{"变量1": "111", "变量2": "222"}},
			want: "您手机注册的验证码为：【111】，如有问题请拨打客服电话：【222】",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMapStrings(tt.args.s, tt.args.news); got != tt.want {
				t.Errorf("ReplaceMapStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceMapStringsForStringIndex(t *testing.T) {
	type args struct {
		s    string
		old  string
		news MapStrings
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "phone",
			args: args{
				s:    "您手机注册的验证码为：【变量】，如有问题请拨打客服电话：【变量】",
				old:  "变量",
				news: MapStrings{"0": "111", "1": "222"},
			},
			want: "您手机注册的验证码为：【111】，如有问题请拨打客服电话：【222】",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMapStringsForStringIndex(tt.args.s, tt.args.old, tt.args.news); got != tt.want {
				t.Errorf("ReplaceMapStringsForStringIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
