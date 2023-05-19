package gosms

import (
	"fmt"
	"strconv"
)

type PhoneNumber struct {
	Number int `json:"number"`
	Code   int `json:"code"`
}

func (p PhoneNumber) I() IPhoneNumber {
	return &p
}

func (p *PhoneNumber) GetCode() int {
	return p.Code
}

func (p *PhoneNumber) GetNumber() int {
	return p.Number
}
func (p *PhoneNumber) GetUniversalNumber() string {
	return fmt.Sprintf("%v%v", p.getPrefixedIDDCode("+"), strconv.Itoa(p.Number))
}
func (p *PhoneNumber) GetZeroPrefixedNumber() string {
	return fmt.Sprintf("%v%v", p.getPrefixedIDDCode("00"), strconv.Itoa(p.Number))
}
func (p *PhoneNumber) getPrefixedIDDCode(prefix string) string {
	if p.Code == 0 {
		return ""
	}
	return fmt.Sprintf("%v%v", prefix, p.Code)
}
