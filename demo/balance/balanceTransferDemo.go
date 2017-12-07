/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package balance

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/balanceTransfer"
)

var TransferDemo = new(BalanceTransferDemo)

type BalanceTransferDemo struct {
	demoAppID      string
	demoTransferId string
}

func (c *BalanceTransferDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *BalanceTransferDemo) Run() {
	balanceTransfer, err := c.New()
	common.Response(balanceTransfer, err)
	c.demoTransferId = balanceTransfer.Id
	common.Response(c.Get())
	common.Response(c.List())
}

func (c *BalanceTransferDemo) New() (*pingpp.BalanceTransfer, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)
	params := &pingpp.BalanceTransferParams{
		Amount:      1,
		User:        "demoUser",
		Recipient:   "demoUser2",
		Order_no:    fmt.Sprintf("%d", orderno),
		Description: "转账",
	}
	return balanceTransfer.New(c.demoAppID, params)
}

//查询 balanceTransfer 对象
func (c *BalanceTransferDemo) Get() (*pingpp.BalanceTransfer, error) {
	return balanceTransfer.Get(c.demoAppID, c.demoTransferId)
}

//查询 BalanceTransfer 对象列表
func (c *BalanceTransferDemo) List() (*pingpp.BalanceTransferList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"

	return balanceTransfer.List(c.demoAppID, params)
}
