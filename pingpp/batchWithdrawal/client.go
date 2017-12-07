package batchWithdrawal

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

func Confirm(appId string, params *pingpp.BatchWithdrawalParams) (*pingpp.BatchWithdrawal, error) {
	return getC().Confirm(appId, params)
}
func (c Client) Confirm(appId string, params *pingpp.BatchWithdrawalParams) (*pingpp.BatchWithdrawal, error) {
	params.Status = "pending"
	paramsString, _ := pingpp.JsonEncode(params)
	batchWithdrawal := &pingpp.BatchWithdrawal{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, nil, paramsString, batchWithdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func Cancel(appId string, params *pingpp.BatchWithdrawalParams) (*pingpp.BatchWithdrawal, error) {
	return getC().Cancel(appId, params)
}
func (c Client) Cancel(appId string, params *pingpp.BatchWithdrawalParams) (*pingpp.BatchWithdrawal, error) {
	params.Status = "canceled"
	paramsString, _ := pingpp.JsonEncode(params)
	batchWithdrawal := &pingpp.BatchWithdrawal{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, nil, paramsString, batchWithdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func Get(appId, batchWithdrawalId string) (*pingpp.BatchWithdrawal, error) {
	return getC().Get(appId, batchWithdrawalId)
}
func (c Client) Get(appId, batchWithdrawalId string) (*pingpp.BatchWithdrawal, error) {
	batchWithdrawal := &pingpp.BatchWithdrawal{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals/%s", appId, batchWithdrawalId), c.Key, nil, nil, batchWithdrawal)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func List(appId string, params *pingpp.PagingParams) (*pingpp.BatchWithdrawalList, error) {
	return getC().List(appId, params)
}
func (c Client) List(appId string, params *pingpp.PagingParams) (*pingpp.BatchWithdrawalList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchWithdrawalList := &pingpp.BatchWithdrawalList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, body, nil, batchWithdrawalList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawalList, err
}
