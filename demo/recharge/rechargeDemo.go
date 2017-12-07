/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package recharge

import (
	"math/rand"
	"time"

	"fmt"

	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/recharge"
)

var Demo = new(RechargeDemo)

type RechargeDemo struct {
	demoAppID      string
	demoRechargeID string
}

func (c *RechargeDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *RechargeDemo) Run() {
	recharge, err := c.New()
	common.Response(recharge, err)
	c.demoRechargeID = recharge.ID
	common.Response(c.Get())
	common.Response(c.List())
}

func (c *RechargeDemo) New() (*pingpp.Recharge, error) {
	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	rechargeParams := &pingpp.RechargeParams{
		User: "1477895856250",
		Charge: pingpp.RechargeCharge{
			Amount:  10,
			Channel: "alipay_qr",
			OrderNo: fmt.Sprintf("%d", orderno),
			Subject: "Go SDK Subject",
			Body:    "Go SDK Body",
		},
		BalanceBonus: pingpp.RechargeBonus{
			Amount: 1,
		},
		Description: "Go SDK Description",
	}
	return recharge.New(c.demoAppID, rechargeParams)
}

func (c *RechargeDemo) Get() (*pingpp.Recharge, error) {
	return recharge.Get(c.demoAppID, c.demoRechargeID)
}

func (c *RechargeDemo) List() (*pingpp.RechargeList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"

	return recharge.List(c.demoAppID, params)
}
