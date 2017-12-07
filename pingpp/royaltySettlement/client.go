package royaltySettlement

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(params *pingpp.RoyaltySettlementCreateParams) (*pingpp.RoyaltySettlement, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.RoyaltySettlementCreateParams) (*pingpp.RoyaltySettlement, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("UserParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	royaltySettlement := &pingpp.RoyaltySettlement{}
	err := c.B.Call("POST", "/royalty_settlements", c.Key, nil, paramsString, royaltySettlement)
	return royaltySettlement, err
}

func Get(royaltySettlementId string) (*pingpp.RoyaltySettlement, error) {
	return getC().Get(royaltySettlementId)
}

func (c Client) Get(royaltySettlementId string) (*pingpp.RoyaltySettlement, error) {
	royaltySettlement := &pingpp.RoyaltySettlement{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), c.Key, nil, nil, royaltySettlement)
	return royaltySettlement, err
}

func Update(royaltySettlementId string, params pingpp.RoyaltySettlementUpdateParams) (*pingpp.RoyaltySettlement, error) {
	return getC().Update(royaltySettlementId, params)
}

func (c Client) Update(royaltySettlementId string, params pingpp.RoyaltySettlementUpdateParams) (*pingpp.RoyaltySettlement, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update RoyaltySettlement  to pingpp is :\n %v\n ", string(paramsString))
	}

	royaltySettlement := &pingpp.RoyaltySettlement{}

	err := c.B.Call("PUT", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), c.Key, nil, paramsString, royaltySettlement)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltySettlement, err
}

func List(params *pingpp.PagingParams) (*pingpp.RoyaltySettlementList, error) {
	return getC().List(params)
}

func (c Client) List(params *pingpp.PagingParams) (*pingpp.RoyaltySettlementList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltySettlementList := &pingpp.RoyaltySettlementList{}
	err := c.B.Call("GET", "/royalty_settlements", c.Key, body, nil, royaltySettlementList)
	return royaltySettlementList, err
}
