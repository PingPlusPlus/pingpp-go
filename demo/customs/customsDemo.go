package customs

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/customs"
)

var Demo = new(CustomsDemo)

type CustomsDemo struct {
	demoAppID string
}

func (c *CustomsDemo) Setup(app string) {
	c.demoAppID = app
}

//创建 customs 对象
func (c *CustomsDemo) New() (*pingpp.Customs, error) {
	params := &pingpp.CustomsParams{
		App:          "app_1Gqj58ynP0mHeX1q",
		Charge:       "ch_L8qn10mLmr1GS8e5OODmHaL4",
		Channel:      "alipay",
		Trade_no:     "12332132131",
		Customs_code: "GUANGZHOU",
		Amount:       8000,
		Is_split:     true,
		Sub_order_no: "123456",
		Extra: map[string]interface{}{
			"pay_account":   "123456",
			"certif_type":   "02",
			"customer_name": "A name",
			"certif_id":     "ID Card No",
			"tax_amount":    10,
		},
	}

	return customs.New(params)
}

//查询 customs 对象
func (c *CustomsDemo) Get() (*pingpp.Customs, error) {
	return customs.Get("14201609281040220109")
}

func (c *CustomsDemo) Run() {
	c.New()
	c.Get()
}
