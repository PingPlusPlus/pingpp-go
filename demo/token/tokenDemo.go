package token

import (
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/token"
)

var Demo = new(TokenDemo)

type TokenDemo struct {
	demoAppID string
}

func (c *TokenDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *TokenDemo) Get() (*pingpp.Token, error) {
	return token.Get("tok_ALeVGaNi5marUpMu4bapUZSZ")
}

func (c *TokenDemo) Run() {
	c.Get()
}
