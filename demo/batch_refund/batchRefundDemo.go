/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package batch_refund

import (
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/batchRefund"
)

var Demo = new(BatchRefundDemo)

type BatchRefundDemo struct {
	demoAppID         string
	charges           []string
	demoBatchRefundID string
}

func (c *BatchRefundDemo) Setup(app string) {
	c.demoAppID = app
	c.charges = []string{"ch_L8qn10mLmr1GS8e5OODmHaL4", "ch_fdOmHaLmLmr1GOD4qn1dS8e5"} // 需要先支付两笔charge，才能做批量退款
}

//创建批量退款
func (c *BatchRefundDemo) New() (*pingpp.BatchRefund, error) {
	params := &pingpp.BatchRefundParams{
		App:         c.demoAppID,
		Batch_no:    "batchrefund" + time.Now().Format("060102150405"),
		Description: "Your Description",
	}

	for _, charge := range c.charges {
		params.Charges = append(params.Charges, map[string]interface{}{
			"charge":      charge,
			"description": "Batch refund description.",
		})
	}
	return batchRefund.New(params)
}

//查询批量退款
func (c *BatchRefundDemo) Get() (*pingpp.BatchRefund, error) {
	return batchRefund.Get(c.demoBatchRefundID)
}

//查询 Batch Refund 对象列表
func (c *BatchRefundDemo) List() (*pingpp.BatchRefundlList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "2")
	params.Filters.AddFilter("app", "", c.demoAppID)
	return batchRefund.List(params)
}

func (c *BatchRefundDemo) Run() {
	batch_refund, err := c.New()
	common.Response(batch_refund, err)
	c.demoBatchRefundID = batch_refund.Id
	batch_refund, err = c.Get()
	common.Response(batch_refund, err)
	batch_refund_list, err := c.List()
	common.Response(batch_refund_list, err)
}
