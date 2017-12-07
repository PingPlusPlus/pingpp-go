/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package royalty

import (
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/royaltyTransaction"
)

var TransactionDemo = new(RoyaltyTransactionDemo)

type RoyaltyTransactionDemo struct {
	demoAppID            string
	royaltyTransactionId string
}

func (c *RoyaltyTransactionDemo) Setup(app string) {
	c.demoAppID = app
	c.royaltyTransactionId = "441170318144700002"
}

// 查询分润明细对象
func (c *RoyaltyTransactionDemo) Get() (*pingpp.RoyaltyTransaction, error) {
	return royaltyTransaction.Get(c.royaltyTransactionId)
}

// 查询分润结算对象列表
func (c *RoyaltyTransactionDemo) List() (*pingpp.RoyaltyTransactionList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "10")
	return royaltyTransaction.List(params)
}

func (c *RoyaltyTransactionDemo) Run() {
	common.Response(c.Get())
	common.Response(c.List())
}
