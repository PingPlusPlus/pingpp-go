package royaltyTransaction

import (
	"fmt"
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

func Get(royaltyTransacitonId string) (*pingpp.RoyaltyTransaction, error) {
	return getC().Get(royaltyTransacitonId)
}

func (c Client) Get(royaltyTransacitonId string) (*pingpp.RoyaltyTransaction, error) {
	royaltyTransaction := &pingpp.RoyaltyTransaction{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_transactions/%s", royaltyTransacitonId), c.Key, nil, nil, royaltyTransaction)
	return royaltyTransaction, err
}

func List(params *pingpp.PagingParams) (*pingpp.RoyaltyTransactionList, error) {
	return getC().List(params)
}

func (c Client) List(params *pingpp.PagingParams) (*pingpp.RoyaltyTransactionList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTransactionList := &pingpp.RoyaltyTransactionList{}
	err := c.B.Call("GET", "/royalty_transactions", c.Key, body, nil, royaltyTransactionList)
	return royaltyTransactionList, err
}
