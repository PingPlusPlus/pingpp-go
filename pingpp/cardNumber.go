package pingpp

type CardNumber struct {
	Object       string `json:"Object"`
	Last4        string `json:"last4"`
	Funding      string `json:"funding"`
	Brand        string `json:"brand"`
	Bank         string `json:"bank"`
	Display_name string `json:"display_name"`
	Logo_url     string `json:"logo_url"`
}

type CardQueryParams struct {
	App         string `json:"app"`
	Card_number string `json:"card_number"`
}
