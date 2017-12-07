package recharge

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

func New(appId string, params *pingpp.RechargeParams) (*pingpp.Recharge, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.RechargeParams) (*pingpp.Recharge, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("RechargeParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of recharge request to pingpp is :\n %v\n ", string(paramsString))
	}

	recharge := &pingpp.Recharge{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/recharges", appId), c.Key, nil, paramsString, recharge)
	return recharge, err
}

func Get(appID, rechargeID string) (*pingpp.Recharge, error) {
	return getC().Get(appID, rechargeID)
}

func (c Client) Get(appID, rechargeID string) (*pingpp.Recharge, error) {
	recharge := &pingpp.Recharge{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s", appID, rechargeID), c.Key, nil, nil, recharge)
	return recharge, err
}

func List(appID string, params *pingpp.PagingParams) (*pingpp.RechargeList, error) {
	return getC().List(appID, params)
}

func (c Client) List(appID string, params *pingpp.PagingParams) (*pingpp.RechargeList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeList := &pingpp.RechargeList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges", appID), c.Key, body, nil, rechargeList)
	return rechargeList, err
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}
