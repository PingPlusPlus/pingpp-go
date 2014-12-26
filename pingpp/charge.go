package pingpp

type ChargeParams struct {
	Order_no  string                 `json:"order_no"`
	App       App                    `json:"app"`
	Channel   string                 `json:"channel"`
	Amount    uint64                 `json:"amount"`
	Currency  string                 `json:"currency"`
	Client_ip string                 `json:"client_ip"`
	Subject   string                 `json:"subject"`
	Body      string                 `json:"body"`
	Metadata  map[string]interface{} `json:"metadata"`
	Extra     Extra                  `json:"extra,omitempty"`
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
	Order_no        string                 `json:"order_no"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Amount          int                    `json:"amount"`
	Amount_settle   uint64                 `json:"amount_settle"`
	Amount_refunded uint64                 `json:"amount_refunded"`
	Time_expire     uint64                 `json:"time_expire"`
	Time_settle     uint64                 `json:"time_settle"`
	Transaction_no  string                 `json:"transaction_no"`
	Currency        string                 `json:"currency"`
	Client_ip       string                 `json:"client_ip"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	Failure_code    int                    `json:"failure_code"`
	Failure_msg     string                 `json:"failure_msg"`
	Extra_data      Extra                  `json:"extra"`
	Metadata        map[string]interface{} `json:"metadata"`
	Refunds         RefundList             `json:"refunds"`
	Credential      interface{}            `json:"credential"`
}

type ChargeList struct {
	Object   string   `json:"object"`
	Url      string   `json:"url"`
	Has_more bool     `json:"has_more"`
	charges  []Charge `json:"charges"`
}

type Extra struct {
	Result_url  string `json:"result_url,omitempty"`
	Success_url string `json:"success_url,omitempty"`
	Cancel_url  string `json:"cancel_url,omitempty"`
	Trade_type  string `json:"trade_type,omitempty"`
	Open_id     string `json:"openid,omitempty"`
	Bfb_login   bool   `json:"bfb_login,omitempty"`
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

// func (charge *Charge) UnmarshalJSON(data []byte) error {
// 	var temp_charge Charge
// 	err := json.Unmarshal(data, &temp_charge)
// 	if err == nil {
// 		*charge = Charge(temp_charge)
// 	}
// 	return err
// }
