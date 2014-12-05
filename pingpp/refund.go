package pingpp

type RefundParams struct {
	Amount      int64
	Description string
}

type RefundListParams struct {
	ChargeID    string
	Limit       int64
	Start_after string
	End_before  string
}

type Refund struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Order_no     string            `json:"order_no"`
	Amount       int64             `json:"amount"`
	Created      uint64            `json:"created"`
	Succeed      bool              `json:"succeed"`
	Time_succeed uint64            `json:"time_succeeded"`
	Description  string            `json:"description"`
	Failure_code int64             `json:"failure_code"`
	Failure_msg  string            `json:"failure_message"`
	Charge_id    string            `json:"charge"`
	Metadata     map[string]string `json:"metadata"`
}

type RefundList struct {
	Object   string   `json:"object"`
	Url      string   `json:"url"`
	Has_more bool     `json:"has_more"`
	Refunds  []Refund `json:"data"`
}

// func (refund *Refund) UnmarshalJson(data []byte) error {
// 	var temp_refund Refund
// 	err := json.Unmarshal(data, &temp_refund)
// 	if err == nil {
// 		*refund = Refund(temp_refund)
// 	} else {
// 		refund.ID = string(data[1 : len(data)-1])
// 	}
// 	return err
// }
