/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package coupon

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/coupon"
)

var Demo = new(CouponDemo)

type CouponDemo struct {
	demoAppID        string
	demoCouponTmplID string
	demoUser         string
}

func (c *CouponDemo) Setup(app string) {
	c.demoAppID = app
	c.demoCouponTmplID = "300216111619300600019101"
	c.demoUser = "uid582d1756b1650"
}

//创建 Coupon 对象
func (c *CouponDemo) New() (*pingpp.Coupon, error) {
	params := &pingpp.CouponParams{
		Coupon_tmpl_id: c.demoCouponTmplID,
	}
	return coupon.New(c.demoAppID, "uid582d1756b1650", params)
}

//批量创建优惠券
func (c *CouponDemo) Batch() (*pingpp.CouponList, error) {
	params := &pingpp.BatchCouponParams{
		Users: []string{"btest1@pingxx.com", "btest2@pingx.com", "btest@pingxx.com"},
	}
	return coupon.BatchNew(c.demoAppID, c.demoCouponTmplID, params)
}

//查询 Coupon 对象
func (c *CouponDemo) Get() (*pingpp.Coupon, error) {
	return coupon.Get(c.demoAppID, c.demoUser, c.demoCouponTmplID)
}

//查询 Coupon 对象列表
func (c *CouponDemo) List() (*pingpp.CouponList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"
	// params.Filters.AddFilter("redeemed", "", "false") //查询用户未核销优惠券

	return coupon.UserList(c.demoAppID, c.demoUser, params)
}

//更新 Coupon 对象
func (c *CouponDemo) Update() (*pingpp.Coupon, error) {
	params := &pingpp.CouponUpdateParams{
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}
	return coupon.Update(c.demoAppID, c.demoUser, c.demoCouponTmplID, params)
}

//删除 Coupon 对象
func (c *CouponDemo) Delete() (*pingpp.DeleteResult, error) {
	return coupon.Delete(c.demoAppID, c.demoUser, c.demoCouponTmplID)
}

func (c *CouponDemo) Run() {
	c.New()
	c.Batch()
	c.Get()
	c.List()
	c.Update()
	c.Delete()
}
