/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package coupon

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/couponTemplate"
)

var TmplDemo = new(CouponTmplDemo)

type CouponTmplDemo struct {
	demoAppID        string
	demoCouponTmplID string
	demoUser         string
}

func (c *CouponTmplDemo) Setup(app string) {
	c.demoAppID = app
	c.demoUser = "uid582d1756b1650"
}

//创建 Coupon template 对象
func (c *CouponTmplDemo) New() (*pingpp.CouponTmpl, error) {
	params := &pingpp.CouponTmplParams{
		Type:       couponTemplate.CASH_COUPON,
		Amount_off: 10,
	}

	return couponTemplate.New(c.demoAppID, params)
}

//查询 Coupon template 对象
func (c *CouponTmplDemo) Get() (*pingpp.CouponTmpl, error) {
	return couponTemplate.Get(c.demoAppID, c.demoCouponTmplID)
}

//更新 Coupon template 对象
func (c *CouponTmplDemo) Update() (*pingpp.CouponTmpl, error) {
	params := &pingpp.CouponTmplUpdateParams{
		Metadata: map[string]interface{}{
			"keys": "value",
		},
	}
	return couponTemplate.Update(c.demoAppID, c.demoCouponTmplID, params)
}

//删除 Coupon template 对象
func (c *CouponTmplDemo) Delete() (*pingpp.DeleteResult, error) {
	return couponTemplate.Delete(c.demoAppID, c.demoCouponTmplID)
}

//查询 Coupon template 对象列表
func (c *CouponTmplDemo) List() (*pingpp.CouponTmplList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"
	return couponTemplate.List(c.demoAppID, params)
}

//查询 Coupon template 模板下的优惠券列表
func (c *CouponTmplDemo) CouponList() (*pingpp.CouponList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "2")
	params.Filters.AddFilter("redeemed", "", "false") //查询用户未核销优惠券

	return couponTemplate.CouponList("app_1Gqj58ynP0mHeX1q", "300216101014204500032301", params)
}

func (c *CouponTmplDemo) Run() {
	c.New()
	c.Get()
	c.Update()
	c.List()
	c.CouponList()
	c.Delete()
}
