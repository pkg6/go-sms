package juhe

type Response struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
	Result    struct {
		Count int    `json:"count"`
		Fee   int    `json:"fee"`
		Sid   string `json:"sid"`
	} `json:"result"`
}
