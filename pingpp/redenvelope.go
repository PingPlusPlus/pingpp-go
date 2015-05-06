package pingpp

type RedEnvelopeParams struct {
	App         *App             `json:"app"`
	Channel     string           `json:"channel"`
	Order_no    string           `json:"order_no"`
	Amount      uint64           `json:"amount"`
	Currency    string           `json:"currency"`
	Recipient   string           `json:"recipient"`
	Subject     string           `json:"subject"`
	Body        string           `json:"body"`
	Description string           `json:"description"`
	Extra       RedEnvelopeExtra `json:"extra"`
}

type RedEnvelopeListParams struct {
	ListParams
	Created int64
}

type RedEnvelope struct {
	Id          string                 `json:"id"`
	Object      string                 `json:"object"`
	Created     uint64                 `json:"created"`
	Livemode    bool                   `json:"livemode"`
	App         string                 `json:"app"`
	Channel     string                 `json:"channel"`
	Order_no    string                 `json:"order_no"`
	Amount      int                    `json:"amount"`
	Currency    string                 `json:"currency"`
	Recipient   string                 `json:"recipient"`
	Subject     string                 `json:"subject"`
	Body        string                 `json:"body"`
	Description string                 `json:"description"`
	Extra       map[string]interface{} `json:"extra"`
}

type RedEnvelopeList struct {
	ListMeta
	RedEnvelopes []*RedEnvelope `json:"redEnvelopes"`
}

type RedEnvelopeExtra struct {
	Nick_name     string `json:"nick_name"`
	Send_name     string `json:"send_name"`
	Logo          string `json:"logo,omitempty"`
	Share_content string `json:"share_content,omitempty"`
	Share_img     string `json:"share_img,omitempty"`
}
