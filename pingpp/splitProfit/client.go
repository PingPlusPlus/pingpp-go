package splitProfit

import (
	"fmt"
	"log"
	"net/url"

	"github.com/pingplusplus/pingpp-go/pingpp"
)

// Client 分账客户端
// 暂时只支持微信渠道特约商户
type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// New 请求分账
func New(params *pingpp.SplitProfitParams) (*pingpp.SplitProfit, error) {
	return getC().New(params)
}

// New 请求分账
func (c Client) New(params *pingpp.SplitProfitParams) (*pingpp.SplitProfit, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("SplitProfitParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create SplitProfitParams is :\n %v\n ", string(paramsString))
	}

	splitProfit := &pingpp.SplitProfit{}
	err := c.B.Call("POST", fmt.Sprintf("/split_profits"), c.Key, nil, paramsString, splitProfit)
	return splitProfit, err
}

// Get 查询分账
func Get(id string) (*pingpp.SplitProfit, error) {
	return getC().Get(id)
}

// Get 查询分账
func (c Client) Get(id string) (*pingpp.SplitProfit, error) {
	splitProfit := &pingpp.SplitProfit{}

	err := c.B.Call("GET", fmt.Sprintf("/split_profits/%s", id), c.Key, nil, nil, splitProfit)
	return splitProfit, err
}

// List 查询分账列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | charge | string |  | optional | 无 | Ping++ 交易成功的 charge ID
// | type | string | optional | 无 | 分账类型: `split_normal` 为普通分账,`split_return` 为完结分账
// | channel | string | [`wx`、`wx_lite`、`wx_pub`、`wx_wap`、`wx_pub_qr`、`wx_pub_scan`] | optional | 无 | 暂时只支持微信渠道
func List(app, charge, typ, channel string, params *pingpp.PagingParams) (pingpp.SplitProfitList, error) {
	return getC().List(app, charge, typ, channel, params)
}

// List 查询分账列表
func (c Client) List(app, charge, typ, channel string, params *pingpp.PagingParams) (pingpp.SplitProfitList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if charge != "" {
		values.Add("charge", charge)
	}
	if typ != "" {
		values.Add("type", typ)
	}
	if channel != "" {
		values.Add("channel", channel)
	}
	params.Filters.AppendTo(values)

	splitProfitList := pingpp.SplitProfitList{}
	err := c.B.Call("GET", "/split_profits", c.Key, values, nil, &splitProfitList)
	return splitProfitList, err
}
