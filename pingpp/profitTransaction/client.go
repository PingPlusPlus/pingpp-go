package profitTransaction

import (
	"fmt"
	"net/url"

	"github.com/pingplusplus/pingpp-go/pingpp"
)

// Client 分账明细客户端
type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// Get 查询分账明细
func Get(id string) (*pingpp.ProfitTransaction, error) {
	return getC().Get(id)
}

// Get 查询分账明细
func (c Client) Get(id string) (*pingpp.ProfitTransaction, error) {
	profitTransaction := &pingpp.ProfitTransaction{}

	err := c.B.Call("GET", fmt.Sprintf("/profit_transactions/%s", id), c.Key, nil, nil, profitTransaction)
	return profitTransaction, err
}

// List 查询分账明细列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | split_profit| string | 17 | optional | 无 | 分账ID
// | split_receiver|  string | 19 | optional | 无 | 分账接收方ID
// | status | string | - | optional | 无 | 分账明细状态
func List(app, splitProfit, splitReceiver, status string, params *pingpp.PagingParams) (pingpp.ProfitTransactionList, error) {
	return getC().List(app, splitProfit, splitReceiver, status, params)
}

// List 查询分账明细列表
func (c Client) List(app, splitProfit, splitReceiver, status string, params *pingpp.PagingParams) (pingpp.ProfitTransactionList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if splitProfit != "" {
		values.Add("split_profit", splitProfit)
	}
	if splitReceiver != "" {
		values.Add("split_receiver", splitReceiver)
	}
	if status != "" {
		values.Add("status", status)
	}
	params.Filters.AppendTo(values)

	profitTransactionList := pingpp.ProfitTransactionList{}
	err := c.B.Call("GET", "/profit_transactions", c.Key, values, nil, &profitTransactionList)
	return profitTransactionList, err
}
