/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package order

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/order"
)

var Demo = new(OrderDemo)

type OrderDemo struct {
	demoAppID    string
	demoUser     string
	demoOrderID  string
	demoChargeID string
	demoChannel  string
}

func (c *OrderDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "alipay_qr"
	c.demoUser = "demo_user"
}

// 创建商品订单
func (c *OrderDemo) New() (*pingpp.Order, error) {
	//针对metadata字段，可以在每一个 order 对象中加入订单的一些详情，如颜色、型号等属性
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//metadata["type"] = "shoes"
	//metadata["size"] = "40"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	params := &pingpp.OrderCreateParams{
		App:               "app_1Gqj58ynP0mHeX1q",
		Uid:               "1477895856250",
		Merchant_order_no: strconv.Itoa(orderno),
		Amount:            20,
		Currency:          "cny",
		Client_ip:         "127.0.0.1",
		Subject:           "Go SDK Subject",
		Body:              "Go SDK Body",
		Description:       "Go SDK Description",
		RoyaltyUsers: []pingpp.RoyaltyUser{
			pingpp.RoyaltyUser{
				User:   "user_test_0001",
				Amount: 10,
			},
			pingpp.RoyaltyUser{
				User:   "user_test_0002",
				Amount: 10,
			},
		},
		//Coupon:"coupon_id"//优惠券Id
		//Actual_amount:900 //使用优惠券后订单实际金额
		Metadata: metadata,
	}

	return order.New(params)
}

// 商品订单支付
func (c *OrderDemo) Pay() (*pingpp.Order, error) {
	var chargeAmount int64 = 1
	orderPayParams := &pingpp.OrderPayParams{
		Charge_amount: &chargeAmount,
		Channel:       c.demoChannel,
		Extra:         common.ChargeExtra[c.demoChannel],
	}
	return order.Pay(c.demoOrderID, orderPayParams)
}

// 商品订单取消
func (c *OrderDemo) Cancel() (*pingpp.Order, error) {
	return order.Cancel(c.demoUser, c.demoOrderID)
}

// 商品订单查询
func (c *OrderDemo) Get() (*pingpp.Order, error) {
	return order.Get(c.demoOrderID)
}

// 商品订单列表查询
func (c *OrderDemo) List() (*pingpp.OrderList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("app", "", c.demoAppID)
	params.Filters.AddFilter("page", "", "1")     //取第一页数据
	params.Filters.AddFilter("per_page", "", "2") //每页两个Order对象
	//params.Filters.AddFilter("created", "", "1475127952")
	params.Filters.AddFilter("paid", "", "true")
	return order.List(params)
}

// 查询 charge
func (c *OrderDemo) Charge() (*pingpp.Charge, error) {
	return order.Charge(c.demoOrderID, c.demoChargeID)
}

// 查询 charge 列表
func (c *OrderDemo) ChargeList() (*pingpp.ChargeList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //取第一页数据
	params.Filters.AddFilter("per_page", "", "2") //每页两个Order对象
	return order.ChargeList(c.demoOrderID, params)
}

// Update 商品订单更新
func (c *OrderDemo) Update() (*pingpp.Order, error) {
	orderUpdateParams := &pingpp.OrderUpdateParams{
		Status: "canceled",
		User:   c.demoUser,
	}
	return order.Update(c.demoOrderID, orderUpdateParams)
}

func (c *OrderDemo) Run() {
	order, err := c.New()
	c.demoOrderID = order.ID
	common.Response(order, err)
	common.Response(c.Update())
	common.Response(c.Pay())
	common.Response(c.Get())
	common.Response(c.List())
	order, _ = c.New()
	c.demoOrderID = order.ID
	common.Response(c.Cancel())
	charges, err := c.ChargeList()
	common.Response(charges, err)
	if len(charges.Values) >= 1 {
		c.demoChargeID = charges.Values[0].ID
		common.Response(c.Charge())
	}
}
