// Package subBank 银行支行
/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package subBank

import (
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/subBank"
)

// Demo 银行支行示例
var Demo = new(demo)

// demo 银行支行示例
type demo struct {
	app string
}

// Setup 设置参数
func (c *demo) Setup(app string) {
	c.app = app
}

// List 查询 银行支行 对象列表
func (c *demo) List() (pingpp.SubBankList, error) {
	return subBank.List(c.app, "0308", "浙江省", "宁波市", "chanpay")
}

// Run 运行示例
func (c *demo) Run() {
	c.List()
}
