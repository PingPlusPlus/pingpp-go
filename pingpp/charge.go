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
	Extra       Extra                  `json:"extra,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Time_expire uint64                 `json:"time_expire,omitempty"`
	Description string                 `json:"description,omitempty"`
}

type ChargeListParams struct {
	Limit       uint64
	Start_after string
	End_before  string
	Createdgt   string "created[gt]"
	Createdgte  string "created[gte]"
	Createdlt   string "created[lt]"
	Createdlte  string "created[lte]"
	Appid       string "app[id]"
	Channel     string
	Paid        uint64
	Refunded    uint64
}

type Charge struct {
	Id              string                 `json:"id"`
	Object          string                 `json:"object"`
	Created         uint64                 `json:"created"`
	Livemode        bool                   `json:"livemode"`
	Paid            bool                   `json:"paid"`
	Refunded        bool                   `json:"refunded"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Order_no        string                 `json:"order_no"`
	Client_ip       string                 `json:"client_ip"`
	Amount          int                    `json:"amount"`
	Amount_settle   uint64                 `json:"amount_settle"`
	Currency        string                 `json:"currency"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	Extra_data      Extra                  `json:"extra"`
	Time_paid       uint64                 `json:"time_paid"`
	Time_expire     uint64                 `json:"time_expire"`
	Time_settle     uint64                 `json:"time_settle"`
	Transaction_no  uint64                 `json:"transaction_no"`
	Refunds         RefundList             `json:"refunds"`
	Amount_refunded uint64                 `json:"amount_refunded"`
	Failure_code    int                    `json:"failure_code"`
	Failure_msg     string                 `json:"failure_msg"`
	Metadata        map[string]interface{} `json:"metadata"`
	Credential      map[string]interface{} `json:"credential"`
	Description     string                 `json:"description"`
}

type ChargeList struct {
	Object   string   `json:"object"`
	Url      string   `json:"url"`
	Has_more bool     `json:"has_more"`
	charges  []Charge `json:"charges"`
}

type Extra struct {
	Result_url    string `json:"result_url,omitempty"`
	Success_url   string `json:"success_url,omitempty"`
	Cancel_url    string `json:"cancel_url,omitempty"`
	Trade_type    string `json:"trade_type,omitempty"`
	Open_id       string `json:"openid,omitempty"`
	Bfb_login     bool   `json:"bfb_login,omitempty"`
	Payment_token string `json:"payment_token,omitempty"`
}

type App struct {
	Id                 string   `json:"id"`
	Object             string   `json:"object,omitempty"`
	Created            uint64   `json:"created,omitempty"`
	Display_name       string   `json:"display_name,omitempty"`
	Notify_url         string   `json:"notify_url,omitempty"`
	Goods_type         int64    `json:"goods_type,omitempty"`
	Channels_supported []string `json:"channels_supported,omitempty"`
}
