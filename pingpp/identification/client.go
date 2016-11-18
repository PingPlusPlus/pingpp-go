package identification

import (
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

const (
	IDENTIFY_IDCARD   = "id_card"
	IDENTIFY_BANKCARD = "bank_card"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func New(params *pingpp.IdentificationParams) (*pingpp.IdentificationResult, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.IdentificationParams) (*pingpp.IdentificationResult, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("IdentificationParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of identification request to pingpp is :\n %v\n ", string(paramsString))
	}
	identificationResult := &pingpp.IdentificationResult{}

	err := c.B.Call("POST", "/identification", c.Key, nil, paramsString, identificationResult)
	return identificationResult, err
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}
