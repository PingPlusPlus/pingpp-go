package pingpp

type TransferParams struct {
	App         App           `json:"app"`
	Channel     string        `json:"channel"`
	Order_no    string        `json:"order_no"`
	Amount      uint64        `json:"amount"`
	Type        string        `json:"type"`
	Currency    string        `json:"currency"`
	Recipient   string        `json:"recipient"`
	Description string        `json:"description"`
	Extra       TransferExtra `json:"extra"`
}

type TransferListParams struct {
	ListParams
	Created int64
}

type Transfer struct {
	Id              string        `json:"id"`
	Object          string        `json:"object"`
	Type            string        `json:"type"`
	Created         int64         `json:"created"`
	Time_transfered int64         `json:"time_transfered"`
	Livemode        bool          `json:"livemode"`
	Status          bool          `json:"status"`
	App             string        `json:"app"`
	Channel         string        `json:"channel"`
	Order_no        string        `json:"order_no"`
	Amount          uint64        `json:"amount"`
	Currency        string        `json:"currency"`
	Recipient       string        `json:"recipient"`
	Description     string        `json:"description"`
	Transaction_no  string        `json:"transaction_no"`
	Extra           TransferExtra `json:"extra"`
}

type TransferExtra struct {
	User_name   string `json:"user_name,omitempty"`
	Force_check bool   `json:"force_check,omitempty"`
}

type TransferList struct {
	ListMeta
	Transfers []*Transfer `json:"transfer"`
}
