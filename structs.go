package gosms

import (
	"reflect"
)

// IsZero 判断Struct是否被初始化
func IsZero(s any) bool {
	valueOf := reflect.ValueOf(s)
	for valueOf.Kind() != reflect.Struct {
		return true
	}
	typeOf := reflect.TypeOf(s)
	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i)
		val := valueOf.FieldByName(field.Name)
		zero := reflect.Zero(val.Type()).Interface()
		current := val.Interface()
		if !reflect.DeepEqual(current, zero) {
			return false
		}
	}
	return true
}
