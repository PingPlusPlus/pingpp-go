/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package charge

import (
	//"encoding/base64"

	"math/rand"
	"strconv"
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	//"io/ioutil"
)

var Demo = new(ChargeDemo)

type ChargeDemo struct {
	demoAppID   string
	demoChannel string
	demoCharge  string
}

func (c *ChargeDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "alipay"
	c.demoCharge = "ch_L8qn10mLmr1GS8e5OODmHaL4"
}

func (c *ChargeDemo) New() (*pingpp.Charge, error) {
	//针对metadata字段，可以在每一个 charge 对象中加入订单的一些详情，如颜色、型号等属性
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//metadata["type"] = "shoes"
	//metadata["size"] = "40"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	params := &pingpp.ChargeParams{
		Order_no:  strconv.Itoa(orderno),
		App:       pingpp.App{Id: c.demoAppID},
		Amount:    1000,
		Channel:   c.demoChannel,
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body",
		Extra:     common.Extra.ChargeExtra[c.demoChannel],
		Metadata:  metadata,
	}

	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	return charge.New(params)
}

// 查询 charge 对象
func (c *ChargeDemo) Get() (*pingpp.Charge, error) {
	return charge.Get(c.demoCharge)
}

// 撤销charge，此接口仅接受线下 isv_scan、isv_wap、isv_qr 渠道的订单调用
func (c *ChargeDemo) Reverse() (*pingpp.Charge, error) {
	return charge.Reverse(c.demoCharge)
}

// 查询 charge 对象列表
func (c *ChargeDemo) List() *charge.Iter {
	params := &pingpp.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return charge.List(c.demoAppID, params)
}

func (c *ChargeDemo) Run() {
	c.New()
	c.Get()
	c.List()
}
