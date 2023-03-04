package gosms

import (
	"errors"
)

// SenderNumbers 多手机号发送短信
func SenderNumbers(numbers []IPhoneNumber, message IMessage, gateway IGateway) (smsResult SMSResults, err error) {
	isSuccessful := false
	for _, number := range numbers {
		result, err := Sender(number, message, gateway)
		if err == nil {
			isSuccessful = true
		}
		smsResult = append(smsResult, result)
	}
	if !isSuccessful {
		return smsResult, errors.New("all the gateways have failed. You can get error details by contracts.ResultsMap.Exception")
	}
	return smsResult, nil
}

// SenderGateways 多渠道发送短信
func SenderGateways(number IPhoneNumber, message IMessage, gateways []IGateway) (smsResult SMSResults, err error) {
	isSuccessful := false
	for _, gateway := range gateways {
		result, err := Sender(number, message, gateway)
		if err == nil {
			isSuccessful = true
		}
		smsResult = append(smsResult, result)
	}
	if !isSuccessful {
		return smsResult, errors.New("all the gateways have failed. You can get error details by contracts.ResultsMap.Exception")
	}
	return smsResult, nil
}

// Sender 通过interface 发送
func Sender(number IPhoneNumber, message IMessage, gateway IGateway) (SMSResult, error) {
	result, err := gateway.Send(number, message)
	if err != nil {
		result.ClientResult.Status = StatusFailure
	} else {
		result.ClientResult.Status = StatusSuccess
	}
	if result.ClientResult.Exception != nil {
		result.ClientResult.Exception = err
	}
	return result, err
}
