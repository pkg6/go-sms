package ihuyi

type Response struct {
	Code  int    `json:"code"`
	SmsId string `json:"smsid"`
	Msg   string `json:"msg"`
}
