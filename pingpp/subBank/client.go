package subBank

import (
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

// Client 支行客户端
type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// List 按银行编号和省市查询支行信息列表
// 参数 | 类型 | 长度/个数/范围 | 是否必须 | 默认值 | 描述
// app | string | 20 | required | 无 | App ID。
// open_bank_code | string | 4 | required | 无 | 银行编号
// prov | string | 1~20 | required | 无 | 省份。
// city | string | 1~40 | required | 无 | 城市。
// channel | string | [`chanpay`] | required | 无 | 渠道。
func List(app, openBankCode, prov, city, channel string) (pingpp.SubBankList, error) {
	return getC().List(app, openBankCode, prov, city, channel)
}

// List 按银行编号和省市查询支行信息列表
func (c Client) List(app, openBankCode, prov, city, channel string) (pingpp.SubBankList, error) {
	values := &url.Values{}
	values.Add("app", app)
	values.Add("open_bank_code", openBankCode)
	values.Add("prov", prov)
	values.Add("city", city)
	values.Add("channel", channel)

	subBankList := pingpp.SubBankList{}
	err := c.B.Call("GET", "/sub_banks", c.Key, values, nil, &subBankList)
	return subBankList, err
}
