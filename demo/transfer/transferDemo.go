/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package transfer

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/transfer"
)

var Demo = new(TransferDemo)

type TransferDemo struct {
	demoAppID   string
	demoChannel string
}

func (c *TransferDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "wx_pub"
}

// 创建 Transfer
func (c *TransferDemo) New() (*pingpp.Transfer, error) {
	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成订单号，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	transferParams := &pingpp.TransferParams{
		App:         pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
		Channel:     c.demoChannel,
		Order_no:    strconv.Itoa(orderno),
		Amount:      100,
		Currency:    "cny",
		Type:        "b2c",
		Recipient:   "youropenid",
		Description: "Your Description",
		Extra:       common.TransferExtra[c.demoChannel],
	}
	return transfer.New(transferParams)
}

// 查询 Transfer
func (c *TransferDemo) Get() (*pingpp.Transfer, error) {
	return transfer.Get("tr_98WLGObXbPKO1avLe5CqH0WH")
}

// 取消 Transfer
func (c *TransferDemo) Reverse() (*pingpp.Transfer, error) {
	return transfer.Reverse("tr_130250408515698647040014")
}

func (c *TransferDemo) List() *transfer.Iter {
	params := &pingpp.TransferListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return transfer.List(params)
}

func (c *TransferDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Reverse()
}
