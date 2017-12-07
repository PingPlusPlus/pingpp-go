/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package channel

import (
	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/channel"
)

var Demo = new(ChannelDemo)

type ChannelDemo struct {
	demoAppID    string
	demoSubAppID string
}

func (c *ChannelDemo) Setup(app string) {
	c.demoAppID = app
	c.demoSubAppID = "app_mXfLCGn5qTa1TOuX"
}

const (
	channelAppId    = "app_1Gqj58ynP0mHeX1q"
	channelSubAppId = "app_mXfLCGn5qTa1TOuX"
	channelName     = "alipay"
)

//创建子商户应用支付渠道
func (c *ChannelDemo) New() (*pingpp.Channel, error) {
	params := &pingpp.ChannelParams{
		Channel: "alipay",
		Params: map[string]interface{}{
			"fee_rate":            60,
			"alipay_pid":          "2088501666666666",
			"alipay_account":      "account@example.com",
			"alipay_security_key": "Your security_key",
			"alipay_mer_app_private_key": `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDSBOW3jdthyqSBMNJ8P+BQnfoKpL29BtvACW1gr8YhXh8EqpBU
nUDdQgi8uYnprXBbR5O1DVnIqLKG9loEn3Rc2iqpnj3M3nSShuVByjyJjQ+DAIG2
/cgJjGQknCLo0CKtuEIyD5xBKYVz3GLofLKqCNGDYdUIxwgaBBpssNIDGQIDAQAB
AoGBAKmzw1taiRawA9VQegRkKQF7ZXwMOjTvwcme1H74CYUU5MOEfzOgDbW7kgvN
cJ8dwlg/sh7uNsppZjif/4UUw5R7bSu33m1sIyglmKUYTU7Kw+ETVAPgwkQjJhek
V/pDr143vmchAblD4RqQZTneojTkvYgci4RkHHHIIZ8lClIBAkEA/nEyCKzl0gxU
LWMd0HKLctcwDu6NPWycffFzSg/+k1+h0GlSTp2E8J6DKOYnrlQYvK2/BnbFPfrb
EySi+7c86QJBANNOExrr7xl54JnlZxbXNDnNrql2brPk1DsV/3Lo3Tmt8NuVqiyo
hVE8Vs/CPRqTTSPoTV4TwSscB4Torlox9rECQB9tne+CY7TJPxCIIKOhsmXR/Kar
gpimtMG9tC7ewOQ1OMiEad06CbSq76p6m0YmLxQHJgRHYV+hf7Pin5sV7BkCQQC6
9KxAuJk/YC9R2r/AXL4vmoU8GLZP4lnIwWjXwaLiwryFfEEp7BywyINCpOgtWED7
UTEK2M2jl9QrSzfgQ66xAkBm2RI+8onm/4PVKtOt8tqLjfsFGMR3g0aUwgSbznc0
Xg9dfU+YUgqfQnyAQHt9jG3/SBdmIrYoWwb7TqJZLkZI
-----END RSA PRIVATE KEY-----`,
			"alipay_app_public_key": `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDSBOW3jdthyqSBMNJ8P+BQnfoK
pL29BtvACW1gr8YhXh8EqpBUnUDdQgi8uYnprXBbR5O1DVnIqLKG9loEn3Rc2iqp
nj3M3nSShuVByjyJjQ+DAIG2/cgJjGQknCLo0CKtuEIyD5xBKYVz3GLofLKqCNGD
YdUIxwgaBBpssNIDGQIDAQAB
-----END PUBLIC KEY-----`,
		},
		Banned:      false,
		BannedMsg:   "",
		Description: "alipay description",
	}
	return channel.New(c.demoAppID, c.demoSubAppID, params)
}

//查询子商户应用支付渠道
func (c *ChannelDemo) Get() (*pingpp.Channel, error) {
	return channel.Get(c.demoAppID, c.demoSubAppID, "alipay")
}

//更新子商户应用支付渠道
func (c *ChannelDemo) Update() (*pingpp.Channel, error) {
	params := pingpp.ChannelUpdateParams{
		Params: map[string]interface{}{
			"hello": "world",
		},
	}
	return channel.Update(channelAppId, channelSubAppId, channelName, params)
}

//删除子商户应用支付渠道
func (c *ChannelDemo) Delete() (*pingpp.ChannelDeleteResult, error) {
	return channel.Delete(channelAppId, channelSubAppId, channelName)
}

func (c *ChannelDemo) Run() {
	channel, err := c.New()
	common.Response(channel, err)
	channel, err = c.Get()
	common.Response(channel, err)
	channel, err = c.Update()
	common.Response(channel, err)
	d, err := c.Delete()
	common.Response(d, err)
}
