package balanceTransaction

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

/*
* 查询用户明细对象
* @param appId string
* @param txnId string
* @return BalanceTransaction
 */
func Get(appId, txnId string) (*pingpp.BalanceTransaction, error) {
	return getC().Get(appId, txnId)
}

func (c Client) Get(appId, txnId string) (*pingpp.BalanceTransaction, error) {
	balanceTransactions := &pingpp.BalanceTransaction{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transactions/%s", appId, txnId), c.Key, nil, nil, balanceTransactions)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get BalanceTransactions error: %v\n", err)
		}
	}
	return balanceTransactions, err
}

/*
* 查询用户明细对象列表
* @param appId string
* @param params PagingParams
* @return BalanceTransactionList
 */
func List(appId string, params *pingpp.PagingParams) (*pingpp.BalanceTransactionList, error) {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *pingpp.PagingParams) (*pingpp.BalanceTransactionList, error) {
	balanceList := &pingpp.BalanceTransactionList{}
	body := &url.Values{}
	params.Filters.AppendTo(body)

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transactions", appId), c.Key, body, nil, balanceList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get BalanceTransactions List error: %v\n", err)
		}
	}
	return balanceList, err
}
