package pingpp

type RedEnvelopeParams struct {
	App         App                    `json:"app"`
	Channel     string                 `json:"channel"`
	Order_no    string                 `json:"order_no"`
	Amount      uint64                 `json:"amount"`
	Currency    string                 `json:"currency"`
	Recipient   string                 `json:"recipient"`
	Subject     string                 `json:"subject"`
	Body        string                 `json:"body"`
	Description string                 `json:"description"`
	Extra       map[string]interface{} `json:"extra"`
}

type RedEnvelopeListParams struct {
	ListParams
	Created int64
}

type RedEnvelope struct {
	Id             string                 `json:"id"`
	Object         string                 `json:"object"`
	Created        uint64                 `json:"created"`
	Received       uint64                 `json:"received"`
	Livemode       bool                   `json:"livemode"`
	Status         string                 `json:"status"`
	App            string                 `json:"app"`
	Channel        string                 `json:"channel"`
	Order_no       string                 `json:"order_no"`
	Transaction_no string                 `json:"transaction_no"`
	Amount         uint64                 `json:"amount"`
	Currency       string                 `json:"currency"`
	Recipient      string                 `json:"recipient"`
	Subject        string                 `json:"subject"`
	Body           string                 `json:"body"`
	Description    string                 `json:"description"`
	Failure_msg     string                 `json:"failure_msg"`
	Extra           map[string]interface{} `json:"extra"`
	Metadata        map[string]interface{} `json:"metadata"`
}

type RedEnvelopeList struct {
	ListMeta
	Values []*RedEnvelope `json:"data"`
}
