package pingpp

type Customer struct {
	ID             string                 `json:"id"`
	Object         string                 `json:"object"`
	Created        int64                  `json:"created"`
	Livemode       bool                   `json:"livemode"`
	App            string                 `json:"app"`
	Email          string                 `json:"email"`
	Currency       string                 `json:"currency"`
	Description    string                 `json:"description"`
	Metadata       map[string]interface{} `json:"metadata"`
	Source         *CardList              `json:"source"`
	Default_source string                 `json:"default_source"`
}

type CustomerParams struct {
	App         string                 `json:"app"`
	Source      interface{}            `json:"source"`
	Description string                 `json:"description,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type CustomerUpdateParams struct {
	Description    string                 `json:"description,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Default_source string                 `json:"default_source,omitempty"`
}

type CustomerList struct {
	ListMeta
	Values []*Customer `json:"data"`
}

type CustomerListParams struct {
	ListParams
	Created int64
}
