/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package batch_withdraw

import (
	"github.com/PingPlusPlus/pingpp-go/pingpp/batchWithdrawal"
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

var Demo = new(BatchWithdraw)

type BatchWithdraw struct {
	demoAppID           string
	demoBatchWithdrawID string
}

func (c *BatchWithdraw) Setup(app string) {
	c.demoAppID = app
}

func (c *BatchWithdraw) Confirm() (*pingpp.BatchWithdrawal, error) {
	params := &pingpp.BatchWithdrawalParams{
		Withdrawals: []string{"1701611150302360654", "1701611151015078981"},
	}
	return batchWithdrawal.Confirm(c.demoAppID, params)
}

func (c *BatchWithdraw) Cancel() (*pingpp.BatchWithdrawal, error) {
	params := &pingpp.BatchWithdrawalParams{
		Withdrawals: []string{"1701611150302360654", "1701611151015078981"},
	}
	return batchWithdrawal.Cancel(c.demoAppID, params)
}

func (c *BatchWithdraw) Get() (*pingpp.BatchWithdrawal, error) {
	return batchWithdrawal.Get(c.demoAppID, "1901611151015122025")
}

func (c *BatchWithdraw) List() (*pingpp.BatchWithdrawalList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("per_page", "", "2")
	//status 参数可选：提现状态，已申请：created，处理中：pending，完成：succeeded，失败：failed，取消：canceled
	params.Filters.AddFilter("status", "", "pending")
	return batchWithdrawal.List(c.demoAppID, params)
}

func (c *BatchWithdraw) Run() {
	bw, err := c.Confirm()
	common.Response(bw, err)
	bw, err = c.Cancel()
	common.Response(bw, err)
	bw, err = c.Get()
	common.Response(bw, err)
	bwlist, err := c.List()
	common.Response(bwlist, err)
}
