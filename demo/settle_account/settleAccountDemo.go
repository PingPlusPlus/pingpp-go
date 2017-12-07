/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package settle_account

import (
	"github.com/PingPlusPlus/pingpp-go/pingpp/settleAccount"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

var Demo = new(SettleAccountDemo)

type SettleAccountDemo struct {
	demoAppID string
}

func (c *SettleAccountDemo) Setup(app string) {
	c.demoAppID = app
}

// 创建结算账户对象-alipay
func (c *SettleAccountDemo) New() (*pingpp.SettleAccount, error) {
	//目前支持的渠道有 bank_account，alipay，wx_pub
	params := &pingpp.SettleAccountParams{
		Channel: "bank_account",
		Recipient: map[string]interface{}{
			"account":        "6214666666666666",
			"name":           "张三",
			"type":           "b2c", //转账类型。b2c：企业向个人付款，b2b：企业向企业付款。
			"open_bank":      "招商银行",
			"open_bank_code": "0308",
		},
	}

	/*
		params := pingpp.SettleAccountParams{
			Channel: "alipay",
			Recipient: map[string]interface{}{
				"account": "account01@alipay.com",
				"name":    "张三",
				"type":    "b2c", //转账类型。b2c：企业向个人付款，b2b：企业向企业付款。
			},
		}

		params := pingpp.SettleAccountParams{
			Channel: "wx_pub",
			Recipient: map[string]interface{}{
				"account":              "open_id",
				"name":                 "张三",
				"type":                 "b2c", //转账类型。b2c：企业向个人付款，b2b：企业向企业付款。
				{"force_check", false}, //是否强制校验收款人姓名。仅当 name 参数不为空时该参数生效。

			},
		}
	*/
	return settleAccount.New("app_1Gqj58ynP0mHeX1q", "test_user_003", params)
}

// 查询结算账户对象
func (c *SettleAccountDemo) Get() (*pingpp.SettleAccount, error) {
	return settleAccount.Get("app_1Gqj58ynP0mHeX1q", "user_004", "320217031816231000001001")
}

// 查询结算账户对象列表
func (c *SettleAccountDemo) List() (*pingpp.SettleAccountList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //取第一页数据
	params.Filters.AddFilter("per_page", "", "2") //每页两个SettleAccount对象
	return settleAccount.List("app_1Gqj58ynP0mHeX1q", "user_004", params)
}

//删除结算账户对象
func (c *SettleAccountDemo) Delete() (*pingpp.DeleteResult, error) {
	return settleAccount.Delete("app_1Gqj58ynP0mHeX1q", "user_004", "320217031816231000001001")
}

func (c *SettleAccountDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Delete()
}
