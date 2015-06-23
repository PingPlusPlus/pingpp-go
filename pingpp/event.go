package pingpp

type Event struct {
	Id               string `json:"id"`
	Created          int64  `json:"created"`
	Livemode         bool   `json:"livemode"`
	Type             string `json:"type"`
	Data             Data   `json:"data"`
	Object           string `json:"object"`
	Pending_webhooks int    `json:"pending_webhooks"`
	Request          string `json:"request"`
}

type Data struct {
	Object Object `json:"object"`
}

type Object struct {
	ID                string                 `json:"id,omitempty"`
	Object            string                 `json:"object,omitempty"`
	Created           int64                  `json:"created,omitempty"`
	Livemode          bool                   `json:"livemode,omitempty"`
	Paid              bool                   `json:"paid,omitempty"`
	Refunded          bool                   `json:"refunded,omitempty"`
	App               string                 `json:"app,omitempty"`
	Channel           string                 `json:"channel,omitempty"`
	Order_no          string                 `json:"order_no,omitempty"`
	Client_ip         string                 `json:"client_ip,omitempty"`
	Amount            uint64                 `json:"amount,omitempty"`
	Amount_settle     uint64                 `json:"amount_settle,omitempty"`
	Currency          string                 `json:"currency,omitempty"`
	Subject           string                 `json:"subject,omitempty"`
	Body              string                 `json:"body,omitempty"`
	Extra             *Extra                 `json:"extra,omitempty"`
	Time_paid         int64                  `json:"time_paid,omitempty"`
	Time_expire       int64                  `json:"time_expire,omitempty"`
	Time_settle       int64                  `json:"time_settle,omitempty"`
	Transaction_no    string                 `json:"transaction_no,omitempty"`
	Refunds           *RefundList            `json:"refunds,omitempty"`
	Amount_refunded   uint64                 `json:"amount_refunded,omitempty"`
	Failure_code      string                 `json:"failure_code,omitempty"`
	Failure_msg       string                 `json:"failure_msg,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	Credential        map[string]interface{} `json:"credential,omitempty"`
	Description       string                 `json:"description,omitempty"`
	Succeed           bool                   `json:"succeed,omitempty"`
	Time_succeed      uint64                 `json:"time_succeed,omitempty"`
	Charge_id         string                 `json:"charge,omitempty"`
	Acct_id           string                 `json:"acct_id,omitempty"`
	App_id            string                 `json:"app_id,omitempty"`
	Acct_display_name string                 `json:"acct_display_name,omitempty"`
	App_display_name  string                 `json:"app_display_name,omitempty"`
	Summary_from      uint64                 `json:"summary_from,omitempty"`
	Summary_to        uint64                 `json:"summary_to,omitempty"`
	Charges_amount    uint64                 `json:"charges_amount,omitempty"`
	Charges_count     uint64                 `json:"charges_count,omitempty"`
}

type EventListParams struct {
	ListParams
	Created int64
}
