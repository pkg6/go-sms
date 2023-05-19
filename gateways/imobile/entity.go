package imobile

type request struct {
	AccountId string `json:"AccountId"`
	AccessKey string `json:"AccessKey"`
	Timestamp int    `json:"Timestamp"`
	Random    int    `json:"Random"`
	ExtendNo  string `json:"ExtendNo"`
	ProductId string `json:"ProductId"`
	PhoneNos  string `json:"PhoneNos"`
	Content   string `json:"Content"`
	SendTime  string `json:"SendTime"`
	OutId     string `json:"OutId"`
}

type Response struct {
	Result     string `json:"Result"`
	Reason     string `json:"Reason"`
	MsgId      int    `json:"MsgId"`
	SplitCount int    `json:"SplitCount"`
}
