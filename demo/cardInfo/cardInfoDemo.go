package card

import (
	"github.com/pingplusplus/pingpp-go/demo/common"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/cardInfo"
)

var Demo = new(CardInfoDemo)

type CardInfoDemo struct {
	demoAppID string
}

func (c *CardInfoDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *CardInfoDemo) New() (*pingpp.CardInfo, error) {
	param := &pingpp.CardInfoParams{App: c.demoAppID, BankAccount: "6228480402564890018"}
	return cardInfo.New(param)
}

func (c *CardInfoDemo) Run() {
	card, err := c.New()
	common.Response(card, err)
}
