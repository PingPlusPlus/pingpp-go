package customs

import (
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(params *pingpp.CustomsParams) (*pingpp.Customs, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.CustomsParams) (*pingpp.Customs, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	customs := &pingpp.Customs{}
	err := c.B.Call("POST", "/customs", c.Key, nil, paramsString, customs)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("New Customs error: %v\n", err)
		}
	}
	return customs, err
}

func Get(Id string) (*pingpp.Customs, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*pingpp.Customs, error) {
	customs := &pingpp.Customs{}
	err := c.B.Call("GET", fmt.Sprintf("/customs/%s", Id), c.Key, nil, nil, customs)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Customs error: %v\n", err)
		}
	}
	return customs, err
}
