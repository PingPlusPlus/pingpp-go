package cardNumber

import (
	"log"
	//"net/url"
	pingpp "pingpp-go/pingpp"
	"time"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

//查询指定 card_number
func Post(params *pingpp.CardQueryParams) (*pingpp.CardNumber, error) {
	return getC().Post(params)
}

func (c Client) Post(params *pingpp.CardQueryParams) (*pingpp.CardNumber, error) {
	start := time.Now()
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("CustomerParams Marshall Errors is : %q\n", errs)
		}
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of card request to pingpp is :\n %v\n ", string(paramsString))
	}

	cardNumber := &pingpp.CardNumber{}
	errch := c.B.Call("POST", "/card_info", c.Key, nil, paramsString, cardNumber)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("Card completed in ", time.Since(start))
	}
	return cardNumber, errch
}
