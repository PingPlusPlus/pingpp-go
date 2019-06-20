// Package splitReceiver 分账接收方
/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package splitReceiver

import (
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/splitReceiver"
)

// Demo 分账接收方示例
var Demo = new(demo)

// demo 分账接收方示例
type demo struct {
	app string
}

// Setup 设置参数
func (c *demo) Setup(app string) {
	c.app = app
}

// Get 查询 分账接收方 对象
func (c *demo) New() (*pingpp.SplitReceiver, error) {
	var params pingpp.SplitReceiverParams
	params.App = c.app
	params.Type = "MERCHANT_ID"  //分账接收方类型
	params.Name = "示例商户全称"       //分账接收方全称
	params.Account = "190001001" //分账接收方帐号
	params.Channel = "wx_pub_qr" //分账接收方使用的渠道
	return splitReceiver.New(&params)
}

// Get 查询 分账接收方 对象
func (c *demo) Get() (*pingpp.SplitReceiver, error) {
	return splitReceiver.Get("recv_1fRbIo0jME7yuL")
}

// List 查询 分账接收方 对象列表
func (c *demo) List() (pingpp.SplitReceiverList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "100")
	return splitReceiver.List(c.app, "", "", params)
}

func (c *demo) Delete() (*pingpp.DeleteResult, error) {
	return splitReceiver.Delete("recv_1fRbIo0jME7yuL")
}

// Run 运行示例
func (c *demo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Delete()
}
