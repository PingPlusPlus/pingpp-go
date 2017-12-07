package card

import (
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/card"
)

var Demo = new(CardDemo)

type CardDemo struct {
	demoAppID string
}

func (c *CardDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *CardDemo) New() (*pingpp.Card, error) {
	cus_id := "cus_ALeWGZ8lsN9Czk"

	sms_code := make(map[string]interface{})
	sms_code["id"] = "sms_BDIQG8JZnQhUN1hAbwYubhP3"
	sms_code["code"] = "123222"
	param := &pingpp.CardParams{
		Source:   "tok_AMBKETCThoW7nYUBgpnvhwfu",
		Sms_code: sms_code,
	}

	return card.New(cus_id, param)
}

func (c *CardDemo) Get() (*pingpp.Card, error) {
	return card.Get("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
}

func (c *CardDemo) List() *card.Iter {
	params := &pingpp.CardListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	cus_id := "cus_ALeWGZ8lsN9Czk"
	return card.List(cus_id, params)
}

func (c *CardDemo) Delete() (map[string]interface{}, error) {
	return card.Delete("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
}

func (c *CardDemo) Run() {
	card, err := c.New()
	common.Response(card, err)
	card, err = c.Get()
	common.Response(card, err)
	i := c.List()
	for i.Next() {
		c := i.Card()
		common.Response(c, nil)
	}
	delete, err := c.Delete()
	common.Response(delete, err)
}
