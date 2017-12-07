package withdrawal

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

func New(appId string, params *pingpp.WithdrawalParams) (*pingpp.Withdrawal, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.WithdrawalParams) (*pingpp.Withdrawal, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	withdrawal := &pingpp.Withdrawal{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/withdrawals", appId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func Get(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	return getC().Get(appId, withdrawalId)
}

func (c Client) Get(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	withdrawal := &pingpp.Withdrawal{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, nil, withdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get BalanceWithdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}

//用户的余额提现列表
func List(appId string, params *pingpp.PagingParams) (*pingpp.WithdrawalList, error) {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *pingpp.PagingParams) (*pingpp.WithdrawalList, error) {
	var body *url.Values
	body = &url.Values{}
	params.Filters.AppendTo(body)
	withdrawalList := &pingpp.WithdrawalList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/withdrawals", appId), c.Key, body, nil, withdrawalList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Withdrawal List error: %v\n", err)
		}
	}
	return withdrawalList, err
}

func Cancel(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	return getC().Cancel(appId, withdrawalId)
}
func (c Client) Cancel(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := pingpp.JsonEncode(cancelParams)

	withdrawal := &pingpp.Withdrawal{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func Confirm(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	return getC().Confirm(appId, withdrawalId)
}
func (c Client) Confirm(appId, withdrawalId string) (*pingpp.Withdrawal, error) {
	confirmParams := struct {
		Status string `json:"status"`
	}{
		Status: "pending",
	}
	paramsString, _ := pingpp.JsonEncode(confirmParams)

	withdrawal := &pingpp.Withdrawal{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}
