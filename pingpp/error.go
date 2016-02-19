package pingpp

const (
	InvalidRequest ErrorType = "invalid_request_error"
	APIErr         ErrorType = "api_error"
)

type (
	// 错误的类型
	ErrorType string
	// 错误的状代码
	ErrorCode string
	// 错误的数据结构
	Error struct {
		Type           ErrorType `json:"type"`
		Msg            string    `json:"message"`
		Code           ErrorCode `json:"code,omitempty"`
		Param          string    `json:"param,omitempty"`
		HTTPStatusCode int       `json:"-"`
	}
)

//返回当前Error数据的json字符串
func (e *Error) Error() string {
	er, _ := JsonEncode(e)
	return string(er)
}
