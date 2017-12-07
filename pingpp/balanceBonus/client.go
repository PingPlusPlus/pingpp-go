package balanceBonus

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

func New(appId string, params *pingpp.BalanceBonusParams) (*pingpp.BalanceBonus, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.BalanceBonusParams) (*pingpp.BalanceBonus, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of balance bonus to pingpp is :\n %v\n ", string(paramsString))
	}
	balanceBonus := &pingpp.BalanceBonus{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/balance_bonuses", appId), c.Key, nil, paramsString, balanceBonus)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance Bonus error: %v\n", err)
		}
	}
	return balanceBonus, err
}

func Get(appId, balanceBonusId string) (*pingpp.BalanceBonus, error) {
	return getC().Get(appId, balanceBonusId)
}

func (c Client) Get(appId, balanceBonusID string) (*pingpp.BalanceBonus, error) {
	balanceBonus := &pingpp.BalanceBonus{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses/%s", appId, balanceBonusID), c.Key, nil, nil, balanceBonus)
	return balanceBonus, err
}

func List(appID string, params *pingpp.PagingParams) (*pingpp.BalanceBonusList, error) {
	return getC().List(appID, params)
}

func (c Client) List(appID string, params *pingpp.PagingParams) (*pingpp.BalanceBonusList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceBonusList := &pingpp.BalanceBonusList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses", appID), c.Key, body, nil, balanceBonusList)
	return balanceBonusList, err
}
