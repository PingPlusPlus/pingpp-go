package settleAccount

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

func New(appId, userId string, params *pingpp.SettleAccountParams) (*pingpp.SettleAccount, error) {
	return getC().New(appId, userId, params)
}

func (c Client) New(appId, userId string, params *pingpp.SettleAccountParams) (*pingpp.SettleAccount, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("SettleAccountParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create SettleAccount is :\n %v\n ", string(paramsString))
	}

	settle_account := &pingpp.SettleAccount{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), c.Key, nil, paramsString, settle_account)
	return settle_account, err
}

func Get(appId, userId, settleAccountId string) (*pingpp.SettleAccount, error) {
	return getC().Get(appId, userId, settleAccountId)
}

func (c Client) Get(appId, userId, settleAccountId string) (*pingpp.SettleAccount, error) {
	settleAccount := &pingpp.SettleAccount{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), c.Key, nil, nil, settleAccount)
	return settleAccount, err
}

func List(appId, userId string, params *pingpp.PagingParams) (*pingpp.SettleAccountList, error) {
	return getC().List(appId, userId, params)
}
func (c Client) List(appId, userId string, params *pingpp.PagingParams) (*pingpp.SettleAccountList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	settleAccountList := &pingpp.SettleAccountList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), c.Key, body, nil, settleAccountList)
	return settleAccountList, err
}

func Delete(appId, userId, settleAccountId string) (*pingpp.DeleteResult, error) {
	return getC().Delete(appId, userId, settleAccountId)
}

func (c Client) Delete(appId, userId, settleAccountId string) (*pingpp.DeleteResult, error) {
	result := &pingpp.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), c.Key, nil, nil, result)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}
