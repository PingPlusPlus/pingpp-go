package pingpp

type ChargeParams struct {
	Order_no    string                 `json:"order_no"`
	App         App                    `json:"app"`
	Channel     string                 `json:"channel"`
	Amount      uint64                 `json:"amount"`
	Currency    string                 `json:"currency"`
	Client_ip   string                 `json:"client_ip"`
	Subject     string                 `json:"subject"`
	Body        string                 `json:"body"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Time_expire uint64                 `json:"time_expire,omitempty"`
	Description string                 `json:"description,omitempty"`
}

type ChargeListParams struct {
	ListParams
	Created int64
}

type Charge struct {
	ID              string                 `json:"id"`
	Object          string                 `json:"object"`
	Created         int64                  `json:"created"`
	Livemode        bool                   `json:"livemode"`
	Paid            bool                   `json:"paid"`
	Refunded        bool                   `json:"refunded"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Order_no        string                 `json:"order_no"`
	Client_ip       string                 `json:"client_ip"`
	Amount          uint64                 `json:"amount"`
	Amount_settle   uint64                 `json:"amount_settle"`
	Currency        string                 `json:"currency"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	Extra           map[string]interface{} `json:"extra"`
	Time_paid       uint64                 `json:"time_paid"`
	Time_expire     uint64                 `json:"time_expire"`
	Time_settle     uint64                 `json:"time_settle"`
	Transaction_no  string                 `json:"transaction_no"`
	Refunds         *RefundList            `json:"refunds"`
	Amount_refunded uint64                 `json:"amount_refunded"`
	Failure_code    string                 `json:"failure_code"`
	Failure_msg     string                 `json:"failure_msg"`
	Metadata        map[string]interface{} `json:"metadata"`
	Credential      map[string]interface{} `json:"credential"`
	Description     string                 `json:"description"`
}

type ChargeList struct {
	ListMeta
	Values []*Charge `json:"data"`
}
