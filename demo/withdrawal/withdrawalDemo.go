/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package withdrawal

import (
	"github.com/PingPlusPlus/pingpp-go/pingpp/withdrawal"
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

var Demo = new(WithdrawalDemo)

type WithdrawalDemo struct {
	demoAppID   string
	demoChannel string
}

func (c *WithdrawalDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "unionpay"
}

// 余额提现申请
func (c *WithdrawalDemo) New() (*pingpp.Withdrawal, error) {
	params := &pingpp.WithdrawalParams{
		User:        "user_001",
		Amount:      20000,
		User_fee:    0,
		Channel:     c.demoChannel,
		Description: "Your description",
		Order_no:    "20160829133002",
		Extra:       common.WithdrawExtra[c.demoChannel],
	}
	return withdrawal.New(c.demoAppID, params)
}

// 余额提现查询
func (c *WithdrawalDemo) Get() (*pingpp.Withdrawal, error) {
	return withdrawal.Get(c.demoAppID, "1701611150302360654")
}

// 余额提现列表查询
func (c *WithdrawalDemo) List() (*pingpp.WithdrawalList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "2")
	return withdrawal.List(c.demoAppID, params)
}

// 余额提现取消
func (c *WithdrawalDemo) Cancel() (*pingpp.Withdrawal, error) {
	return withdrawal.Cancel(c.demoAppID, "1701611150302360654")
}

// 余额提现确认
func (c *WithdrawalDemo) Confirm() (*pingpp.Withdrawal, error) {
	return withdrawal.Confirm(c.demoAppID, "1701611150302360654")
}

func (c *WithdrawalDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Cancel()
	c.Confirm()
}
