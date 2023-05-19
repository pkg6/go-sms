package gosms

import (
	"fmt"
	"strconv"
)

// GetPhoneNumber 获取手机号码
func GetPhoneNumber(number IPhoneNumber) string {
	var phoneNumber string
	if number.GetCode() != 0 {
		phoneNumber = fmt.Sprintf("%v%v", number.GetCode(), strconv.Itoa(number.GetNumber()))
	} else {
		phoneNumber = strconv.Itoa(number.GetNumber())
	}
	return phoneNumber
}

// NoCodePhoneNumber 无code手机号
func NoCodePhoneNumber(number int) IPhoneNumber {
	return PhoneNumber{Number: number}.I()
}

// CHNPhoneNumber 中国手机号
func CHNPhoneNumber(number int) IPhoneNumber {
	return PhoneNumber{Number: number, Code: 86}.I()
}
