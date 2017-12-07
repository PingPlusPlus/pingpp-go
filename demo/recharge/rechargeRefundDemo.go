/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package recharge

import (
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/recharegRefund"
)

var RefundDemo = new(RechargeRefundDemo)

type RechargeRefundDemo struct {
	demoAppID      string
	demoRechargeID string
	demoRefundID   string
}

func (c *RechargeRefundDemo) Setup(app string) {
	c.demoRechargeID = "RECHARGE ID"
	c.demoAppID = app
}

func (c *RechargeRefundDemo) Run() {
	refund, err := c.New()
	common.Response(refund, err)
	c.demoRefundID = refund.ID
	common.Response(c.Get())
	common.Response(c.List())
}

func (c *RechargeRefundDemo) New() (*pingpp.Refund, error) {
	rechargeParams := &pingpp.RechargeRefundParams{
		Description: "充值退款",
	}
	return recharegRefund.New(c.demoAppID, c.demoRechargeID, rechargeParams)
}

func (c *RechargeRefundDemo) Get() (*pingpp.Refund, error) {
	return recharegRefund.Get(c.demoAppID, c.demoRechargeID, c.demoRefundID)
}

func (c *RechargeRefundDemo) List() (*pingpp.RefundList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"

	return recharegRefund.List(c.demoAppID, c.demoRechargeID, params)
}
