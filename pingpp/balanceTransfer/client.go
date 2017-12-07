package balanceTransfer

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

func New(appId string, params *pingpp.BalanceTransferParams) (*pingpp.BalanceTransfer, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.BalanceTransferParams) (*pingpp.BalanceTransfer, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of balance transfer to pingpp is :\n %v\n ", string(paramsString))
	}
	balanceTransfer := &pingpp.BalanceTransfer{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/balance_transfers", appId), c.Key, nil, paramsString, balanceTransfer)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance Transfer error: %v\n", err)
		}
	}
	return balanceTransfer, err
}

func Get(appId, balanceTransferId string) (*pingpp.BalanceTransfer, error) {
	return getC().Get(appId, balanceTransferId)
}

func (c Client) Get(appId, balanceTransferID string) (*pingpp.BalanceTransfer, error) {
	balanceTransfer := &pingpp.BalanceTransfer{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers/%s", appId, balanceTransferID), c.Key, nil, nil, balanceTransfer)
	return balanceTransfer, err
}

func List(orderId string, params *pingpp.PagingParams) (*pingpp.BalanceTransferList, error) {
	return getC().List(orderId, params)
}

func (c Client) List(appID string, params *pingpp.PagingParams) (*pingpp.BalanceTransferList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceTransferList := &pingpp.BalanceTransferList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers", appID), c.Key, body, nil, balanceTransferList)
	return balanceTransferList, err
}
