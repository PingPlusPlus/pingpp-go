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
	"github.com/pingplusplus/pingpp-go/pingpp/balanceBonus"
)

var BonusDemo = new(BalanceBonusDemo)

type BalanceBonusDemo struct {
	demoAppID   string
	demoBonusId string
}

func (c *BalanceBonusDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *BalanceBonusDemo) Run() {
	bonus, err := c.New()
	common.Response(bonus, err)
	c.demoBonusId = bonus.Id
	bonus, err = c.Get()
	common.Response(bonus, err)
	bonusList, err := c.List()
	common.Response(bonusList, err)
}

func (c *BalanceBonusDemo) New() (*pingpp.BalanceBonus, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)
	params := &pingpp.BalanceBonusParams{
		Amount:      1,
		User:        "demoUser",
		Order_no:    fmt.Sprintf("%d", orderno),
		Description: "赠送金额",
	}
	return balanceBonus.New(c.demoAppID, params)
}

//查询 balanceBonus 对象
func (c *BalanceBonusDemo) Get() (*pingpp.BalanceBonus, error) {
	return balanceBonus.Get(c.demoAppID, c.demoBonusId)
}

//查询 balanceBonus 对象列表
func (c *BalanceBonusDemo) List() (*pingpp.BalanceBonusList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"

	return balanceBonus.List(c.demoAppID, params)
}
