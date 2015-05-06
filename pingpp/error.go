package pingpp

import "encoding/json"

type ErrorType string

type ErrorCode string

const (
	InvalidRequest ErrorType = "invalid_request_error"
	APIErr         ErrorType = "api_error"
)

// Error is the response returned when a call is unsuccessful.
type Error struct {
	Type           ErrorType `json:"type"`
	Msg            string    `json:"message"`
	Code           ErrorCode `json:"code,omitempty"`
	Param          string    `json:"param,omitempty"`
	HTTPStatusCode int       `json:"-"`
}

// Error serializes the Error object and prints the JSON string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
