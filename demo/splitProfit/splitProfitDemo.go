// Package splitProfit 分账
/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package splitProfit

import (
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/splitProfit"
)

// Demo 分账示例
var Demo = new(demo)

// demo 分账示例
type demo struct {
	app string
}

// Setup 设置参数
func (c *demo) Setup(app string) {
	c.app = app
}

// Get 查询 分账 对象
func (c *demo) New() (*pingpp.SplitProfit, error) {
	var params pingpp.SplitProfitParams
	params.App = c.app
	params.Charge = "ch_aDC44OKyL8yHPG0yX9yzLy5K"
	params.OrderNo = "order_no" // 分账单号，由商家自行生成，规则参照微信分账参数规则
	params.Type = "split_normal"
	params.Recipients = []pingpp.SplitProfitRecipientParams{
		{
			SplitReceiver: "recv_1fRbIo5YgIM4hl",
			Amount:        6,
			Name:          "示例商户全称",
			Description:   "Your Description",
		},
	}
	return splitProfit.New(&params)
}

// Get 查询 分账 对象
func (c *demo) Get() (*pingpp.SplitProfit, error) {
	return splitProfit.Get("sp_1iXmpPJLe71sf9")
}

// List 查询 分账 对象列表
func (c *demo) List() (pingpp.SplitProfitList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "100")
	return splitProfit.List(c.app, "", "", "", params)
}

// Run 运行示例
func (c *demo) Run() {
	c.New()
	c.Get()
	c.List()
}
