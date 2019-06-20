package splitReceiver

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

// New 添加分账接收方
// ##### 微信分账接收方类型 type 说明
// - `MERCHANT_ID`：商户ID
// - `PERSONAL_WECHATID`：个人微信号
// - `PERSONAL_OPENID`：个人openid（由父商户APPID转换得到）
// - `PERSONAL_SUB_OPENID`: 个人sub_openid（由子商户APPID转换得到）
// ##### 微信分账接收方帐号 account 说明
// - 微信分账接收方类型是`MERCHANT_ID`时，是商户ID
// - 微信分账接收方类型是`PERSONAL_WECHATID`时，是个人微信号
// - 微信分账接收方类型是`PERSONAL_OPENID`时，是个人`openid`
// - 微信分账接收方类型是`PERSONAL_SUB_OPENID`时，是个人`sub_openid`
func New(params *pingpp.SplitReceiverParams) (*pingpp.SplitReceiver, error) {
	return getC().New(params)
}

// New 添加分账接收方
func (c Client) New(params *pingpp.SplitReceiverParams) (*pingpp.SplitReceiver, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("SplitReceiverParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create SplitReceiverParams is :\n %v\n ", string(paramsString))
	}

	splitReceiver := &pingpp.SplitReceiver{}
	err := c.B.Call("POST", fmt.Sprintf("/split_receivers"), c.Key, nil, paramsString, splitReceiver)
	return splitReceiver, err
}

// Get 查询分账接收方
func Get(id string) (*pingpp.SplitReceiver, error) {
	return getC().Get(id)
}

// Get 查询分账接收方
func (c Client) Get(id string) (*pingpp.SplitReceiver, error) {
	splitReceiver := &pingpp.SplitReceiver{}

	err := c.B.Call("GET", fmt.Sprintf("/split_receivers/%s", id), c.Key, nil, nil, splitReceiver)
	return splitReceiver, err
}

// List 查询分账接收方列表
// | 参数 | 类型 | 长度/个数/范围 | 是否必需 | 默认值 | 说明
// | --- | --- | --- | --- | --- | ---
// | app | string | 20 | required | 无 | App ID。
// | type | string | [1~32] | optional | 无 | 分账接收方类型
// | channel | string | [`wx`、`wx_lite`、`wx_pub`、`wx_wap`、`wx_pub_qr`、`wx_pub_scan`] | optional | 无 | 暂时只支持微信渠道
func List(app, typ, channel string, params *pingpp.PagingParams) (pingpp.SplitReceiverList, error) {
	return getC().List(app, typ, channel, params)
}

// List 查询分账接收方列表
func (c Client) List(app, typ, channel string, params *pingpp.PagingParams) (pingpp.SplitReceiverList, error) {
	values := &url.Values{}
	values.Add("app", app)
	if typ != "" {
		values.Add("type", typ)
	}
	if channel != "" {
		values.Add("channel", channel)
	}
	params.Filters.AppendTo(values)

	splitReceiverList := pingpp.SplitReceiverList{}
	err := c.B.Call("GET", "/split_receivers", c.Key, values, nil, &splitReceiverList)
	return splitReceiverList, err
}

// Delete 删除分账接收方
func Delete(id string) (*pingpp.DeleteResult, error) {
	return getC().Delete(id)
}

// Delete 删除分账接收方
func (c Client) Delete(id string) (*pingpp.DeleteResult, error) {
	deleteResult := &pingpp.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/split_receivers/%s", id), c.Key, nil, nil, deleteResult)
	return deleteResult, err
}
