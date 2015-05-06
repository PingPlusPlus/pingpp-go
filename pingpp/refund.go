package pingpp

import (
	"encoding/json"
)

type RefundParams struct {
	Amount      uint64                 `json:"amount"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// RefundListParams is the set of parameters that can be used when listing refunds.
type RefundListParams struct {
	ListParams
	Charge string
}

type Refund struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Order_no     string            `json:"order_no"`
	Amount       uint64            `json:"amount"`
	Succeed      bool              `json:"succeed"`
	Created      uint64            `json:"created"`
	Time_succeed uint64            `json:"time_succeed"`
	Description  string            `json:"description"`
	Failure_code string            `json:"failure_code"`
	Failure_msg  string            `json:"failure_msg"`
	Metadata     map[string]string `json:"metadata"`
	Charge_id    string            `json:"charge"`
}

type RefundList struct {
	ListMeta
	Values []*Refund `json:"data"`
}

// UnmarshalJSON handles deserialization of a Refund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Refund) UnmarshalJSON(data []byte) error {
	type refund Refund
	var rr refund
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Refund(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}
	return nil
}
