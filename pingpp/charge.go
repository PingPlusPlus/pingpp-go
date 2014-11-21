package pingpp

type ChargeParams struct {
	order_no  string
	appid     string
	channel   string
	amount    uint64
	currency  string
	client_ip string
	subject   string
	body      string
	metadata  map[string]string
	Extra     *Extra `json:"extra"`
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
	Id              string
	Object          string            `json:"object"`
	Created         string            `json:"created"`
	Livemode        bool              `json:"livemode"`
	Paid            bool              `json:"paid"`
	Refunded        bool              `json:"refunded"`
	Order_no        string            `json:"order_no"`
	App             string            `json:"app"`
	Channel         string            `json:"channel"`
	Amount          uint64            `json:"amount"`
	Amount_settle   uint64            `json:"amount_settle"`
	Amount_refunded uint64            `json:"amount_refunded"`
	Time_expire     string            `json:"time_expire"`
	Time_settle     string            `json:"time_settle"`
	Transaction_no  string            `json:"transaction_no"`
	Currency        string            `json:"currency"`
	Client_ip       string            `json:"client_ip"`
	Subject         string            `json:"subject"`
	Body            string            `json:"body"`
	Failure_code    int               `json:"failure_code"`
	Failure_msg     string            `json:"failure_msg"`
	Metadata        map[string]string `json:"metadata"`
	Refunds         *RefundList       `json:"refunds"`
	Credential      *Credential       `json:"credential"`
}

type ChargeList struct {
	Object   string
	Url      string
	Has_more bool
	charges  []*Charge
}

type Extra struct {
	Result_url  string `json:"result_url"`
	Success_url string `json:"success_url"`
	Cancel_url  string `json:"cancel_url"`
}

// func (charge *Charge) UnmarshalJSON(data []byte) error {
// 	var temp_charge Charge
// 	err := json.Unmarshal(data, &temp_charge)
// 	if err == nil {
// 		*charge = Charge(temp_charge)
// 	}
// 	return err
// }
