package pingpp

type SmsCode struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Created   int64  `json:"created"`
	Validated bool   `json:"validated"`
	Source    string `json:"source"`
	Code      string `json:"code"`
}
