package aliyun

type Response struct {
	RequestID string `json:"RequestId"`
	HostID    string `json:"HostId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
}
