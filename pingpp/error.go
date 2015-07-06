package pingpp

type ErrorType string

type ErrorCode string

const (
	InvalidRequest ErrorType = "invalid_request_error"
	APIErr         ErrorType = "api_error"
)

type Error struct {
	Type           ErrorType `json:"type"`
	Msg            string    `json:"message"`
	Code           ErrorCode `json:"code,omitempty"`
	Param          string    `json:"param,omitempty"`
	HTTPStatusCode int       `json:"-"`
}

func (e *Error) Error() string {
	er, _ := JsonEncode(e)
	return string(er)
}
