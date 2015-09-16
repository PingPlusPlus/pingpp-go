package pingpp

type Card struct {
	ID       string `json:"id"`
	Object   string `json:"object"`
	Created  int64  `json:"created"`
	Last4    string `json:"last4"`
	Funding  string `json:"funding"`
	Brand    string `json:"brand"`
	Bank     string `json:"bank"`
	Customer string `json:"customer"`
}

type CardParams struct {
	Source interface{} `json:"source"`
}

type CardList struct {
	ListMeta
	Values []*Card `json:"data"`
}

type CardListParams struct {
	ListParams
	Created int64
}
