package pingpp

type TransferParams struct {
	App         App                    `json:"app"`
	Channel     string                 `json:"channel"`
	Order_no    string                 `json:"order_no"`
	Amount      uint64                 `json:"amount"`
	Type        string                 `json:"type"`
	Currency    string                 `json:"currency"`
	Recipient   string                 `json:"recipient"`
	Description string                 `json:"description"`
	Extra       map[string]interface{} `json:"extra"`
}

type TransferListParams struct {
	ListParams
	Created int64
}

type Transfer struct {
	Id              string                 `json:"id"`
	Object          string                 `json:"object"`
	Type            string                 `json:"type"`
	Created         int64                  `json:"created"`
	Time_transfered int64                  `json:"time_transfered"`
	Livemode        bool                   `json:"livemode"`
	Status          string                 `json:"status"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Order_no        string                 `json:"order_no"`
	Amount          uint64                 `json:"amount"`
	Currency        string                 `json:"currency"`
	Recipient       string                 `json:"recipient"`
	Description     string                 `json:"description"`
	Transaction_no  string                 `json:"transaction_no"`
	Extra           map[string]interface{} `json:"extra"`
}

type TransferList struct {
	ListMeta
	Values []*Transfer `json:"data"`
}
