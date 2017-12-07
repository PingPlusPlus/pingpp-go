package red_envelope

import (
	"math/rand"
	"strconv"
	"time"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/redEnvelope"
)

var Demo = new(RedEnvelopeDemo)

type RedEnvelopeDemo struct {
	demoAppID string
}

func (c *RedEnvelopeDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *RedEnvelopeDemo) New() (*pingpp.RedEnvelope, error) {
	extra := make(map[string]interface{})
	extra["nick_name"] = "Nick Name"
	extra["send_name"] = "Send Name"
	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	redenvelopeParams := &pingpp.RedEnvelopeParams{
		App:         pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
		Channel:     "wx_pub",
		Order_no:    strconv.Itoa(orderno),
		Amount:      100,
		Currency:    "cny",
		Recipient:   "youropenid",
		Subject:     "Your Subject",
		Body:        "Your Body",
		Description: "Your Description",
		Extra:       extra,
	}
	return redEnvelope.New(redenvelopeParams)
}

func (c *RedEnvelopeDemo) Get() (*pingpp.RedEnvelope, error) {
	return redEnvelope.Get("red_id")
}

func (c *RedEnvelopeDemo) List() *redEnvelope.Iter {
	params := &pingpp.RedEnvelopeListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return redEnvelope.List(params)
}

func (c *RedEnvelopeDemo) Run() {
	c.New()
	c.Get()
	c.List()
}
