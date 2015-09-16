package pingpp

type Token struct {
	ID        string                 `json:"id"`
	Object    string                 `json:"object"`
	Created   int64                  `json:"created"`
	Livemode  bool                   `json:"livemode"`
	Used      bool                   `json:"used"`
	Time_used int64                  `json:"time_used"`
	Type      string                 `json:"type"`
	Card      map[string]interface{} `json:"card"`
	Cms_code  map[string]interface{} `json:"sms_code"`
}

type TokenParams struct {
	Order_no string      `json:"order_no"`
	Amount   uint64      `json:"amount"`
	App      string      `json:"app"`
	Card     interface{} `json:"card"`
}
