/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package royalty

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/royaltySettlement"
)

var SettlementDemo = new(RoyaltySettlementDemo)

type RoyaltySettlementDemo struct {
	demoAppID           string
	royaltySettlementId string
}

func (c *RoyaltySettlementDemo) Setup(app string) {
	c.demoAppID = app
	c.royaltySettlementId = "431170320181400001"
}

//创建分润结算对象
func (c *RoyaltySettlementDemo) New() (*pingpp.RoyaltySettlement, error) {
	params := &pingpp.RoyaltySettlementCreateParams{
		PayerApp:     c.demoAppID,
		Method:       "alipay",
		RecipientApp: c.demoAppID,
		Created: pingpp.Created{
			GT: 1489826451,
			LT: 1492418451,
		},
		SourceUser: "user_002",
		MinAmount:  1,
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}
	return royaltySettlement.New(params)
}

// 查询分润结算对象
func (c *RoyaltySettlementDemo) Get() (*pingpp.RoyaltySettlement, error) {
	return royaltySettlement.Get(c.royaltySettlementId)
}

// 查询分润结算对象列表
func (c *RoyaltySettlementDemo) List() (*pingpp.RoyaltySettlementList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("per_page", "", "3")
	params.Filters.AddFilter("payer_app", "", "app_1Gqj58ynP0mHeX1q")
	return royaltySettlement.List(params)
}

//更新分润结算对象
func (c *RoyaltySettlementDemo) Update() (*pingpp.RoyaltySettlement, error) {
	params := pingpp.RoyaltySettlementUpdateParams{
		Status: "pending",
	}
	return royaltySettlement.Update(c.royaltySettlementId, params)
}

func (c *RoyaltySettlementDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Update()
}
