package identification

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/identification"
)

var Demo = new(IdentificationDemo)

type IdentificationDemo struct {
	demoAppID string
}

func (c *IdentificationDemo) Setup(app string) {
	c.demoAppID = app
}

// 调用身份证认证接口
func (c *IdentificationDemo) New() (*pingpp.IdentificationResult, error) {
	identificationData := map[string]interface{}{
		"id_name":   "张三",
		"id_number": "310181198910107641",
	}
	params := &pingpp.IdentificationParams{
		Type: identification.IDENTIFY_IDCARD, //IDENTIFY_IDCARD:身份证认证;IDENTIFY_BANKCARD:银行卡信息认证
		App:  c.demoAppID,
		Data: identificationData,
	}

	return identification.New(params)
}

func (c *IdentificationDemo) Run() {
	c.New()
}
