package pingpp

type RefundParams struct {
	Amount      uint64                 `json:"amount"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type RefundListParams struct {
	ListParams
	Charge string
}

type Refund struct {
	ID           string                 `json:"id"`
	Object       string                 `json:"object"`
	Order_no     string                 `json:"order_no"`
	Amount       uint64                 `json:"amount"`
	Succeed      bool                   `json:"succeed"`
	Status       string                 `json:"status"`
	Created      uint64                 `json:"created"`
	Time_succeed uint64                 `json:"time_succeed"`
	Description  string                 `json:"description"`
	Failure_code string                 `json:"failure_code"`
	Failure_msg  string                 `json:"failure_msg"`
	Metadata     map[string]interface{} `json:"metadata"`
	Charge_id    string                 `json:"charge"`
}

type RefundList struct {
	ListMeta
	Values []*Refund `json:"data"`
}
