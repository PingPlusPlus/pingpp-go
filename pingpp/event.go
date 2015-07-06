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
	Object map[string]interface{} `json:"object"`
}

type Summary struct {
	Acct_id           string `json:"acct_id,omitempty"`
	App_id            string `json:"app_id.omitempty"`
	Acct_display_name string `json:"acct_display_name"`
	App_display_name  string `json:"app_display_name"`
	Summary_from      uint64 `json:"summary_from"`
	Summary_to        uint64 `json:"summary_to"`
	Charges_amount    uint64 `json:"charges_amount"`
	Charges_count     uint64 `json:"charges_count"`
}

type EventListParams struct {
	ListParams
	Created int64
}

type EventList struct {
	ListMeta
	Values []*Event `json:"data"`
}
