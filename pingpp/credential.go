package pingpp

// import "encoding/json"

type Credential struct {
	object           string      `json:"object"`
	WxCredential     interface{} `json:"wx"`
	AlipayCredential interface{} `json:"alipay"`
	UpmpCredential   interface{} `json:"upmp"`
}

type WxCredential struct {
	appid        string `json:"appId"`
	partnerId    string `json:"partnerId"`
	prepayId     string `json:"prepayId"`
	nonceStr     string `json:"nonceStr"`
	timeStamp    int64  `json:"timeStamp"`
	packageValue string `json:"packageValue"`
	sign         string `json:"sign"`
}

type AlipayCredential struct {
	alipayOrderInfo string `json:"orderInfo"`
}

type UpmpCredential struct {
	tn   string `json:"tn"`
	mode string `json:"mode"`
}
