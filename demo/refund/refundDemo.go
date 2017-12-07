package refund

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/refund"
)

var Demo = new(RefundDemo)

type RefundDemo struct {
	demoAppID string
}

func (c *RefundDemo) Setup(app string) {
	c.demoAppID = app
}

// 通过发起一次退款请求创建一个新的 refund 对象，只能对已经发生交易并且没有全额退款的 charge 对象发起退款
func (c *RefundDemo) New() (*pingpp.Refund, error) {
	params := &pingpp.RefundParams{
		Amount:      1, //可以注释不上传
		Description: "12345",
	}
	return refund.New("ch_id", params) //ch_id 是已付款的订单号
}

// 查询退款对象
func (c *RefundDemo) Get() (*pingpp.Refund, error) {
	return refund.Get("ch_id", "refund_id")
}

// 查询退款列表
func (c *RefundDemo) List() *refund.Iter {
	params := &pingpp.RefundListParams{}
	// params.Filters.AddFilter("limit", "", "1")
	// //设置是不是只需要之前设置的 limit 这一个查询参数
	// params.Single = true
	return refund.List("ch_id", params)
}

func (c *RefundDemo) Run() {
	c.New()
	c.Get()
	c.List()
}
