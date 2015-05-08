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
	Created uint64
}

type RedEnvelope struct {
	Id          string                 `json:"id"`
	Object      string                 `json:"object"`
	Created     uint64                 `json:"created"`
	Livemode    bool                   `json:"livemode"`
	App         string                 `json:"app"`
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

type RedEnvelopeList struct {
	ListMeta
	RedEnvelopes []*RedEnvelope `json:"redEnvelopes"`
}

type RedEnvelopeExtra struct {
	Nick_name     string `json:"nick_name"`
	Send_name     string `json:"send_name"`
	Logo          string `json:"logo,omitempty"`
	Share_url     string `json:"share_url.omitempty"`
	Share_content string `json:"share_content,omitempty"`
	Share_img     string `json:"share_img,omitempty"`
}

func (r *RedEnvelope) UnmarshalJSON(data []byte) error {
	type redEnvelope RedEnvelope
	var rr redEnvelope
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Refund(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}
	return nil
}
