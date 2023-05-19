package gosms

const StatusSuccess = "success"
const StatusFailure = "failure"

// SMSResults 批量发
type SMSResults []SMSResult

// SMSResult 网关统一的结构体
type SMSResult struct {
	Gateway      IGateway     `json:"gateway"`
	PhoneNumber  IPhoneNumber `json:"phone_number"`
	Message      IMessage     `json:"message"`
	ClientResult ClientResult `json:"client_result"`
}

// ClientResult 网关响应结果的结构体
type ClientResult struct {
	// 状态值
	Status string `json:"status" xml:"status"`
	// 请求返回的数据
	Response any `json:"response" xml:"response"`
	// 错误返回的异常
	Exception error `json:"exception" xml:"exception"`
}

// BuildSMSResult 在网关快速生成结构体
func BuildSMSResult(phoneNumber IPhoneNumber, message IMessage, gateway IGateway, response any) SMSResult {
	rt := SMSResult{
		Gateway:     gateway,
		Message:     message,
		PhoneNumber: phoneNumber,
	}
	rt.ClientResult.Response = response
	return rt
}
