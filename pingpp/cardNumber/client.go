package cardNumber

import (
	"fmt"
	"log"
	"net/url"
	pingpp "pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

//查询指定 card_number
func Get(cus_num string, params *pingpp.CardQueryParams) (*pingpp.CardNumber, error) {
	return getC().Get(cus_num, params)
}

func (c Client) Get(cus_num string, params *pingpp.CardQueryParams) (*pingpp.CardNumber, error) {
	var body *url.Values
	if params != nil {
		body = &url.Values{}

		if params.App != "" {
			body.Add("app", params.App)
		}

	}
	cardNumber := &pingpp.CardNumber{}
	err := c.B.Call("GET", fmt.Sprintf("/card_number/%v", cus_num), c.Key, body, nil, cardNumber)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get CardNumber error: %v\n", err)
		}
	}
	return cardNumber, err
}
