package pingpp

// import "encoding/json"

type Credential struct {
	Object           string           `json:"object"`
	WxCredential     WxCredential     `json:"wx"`
	AlipayCredential AlipayCredential `json:"alipay"`
	UpmpCredential   UpmpCredential   `json:"upmp"`
}

type WxCredential struct {
	Appid        string `json:"appId"`
	PartnerId    string `json:"partnerId"`
	PrepayId     string `json:"prepayId"`
	NonceStr     string `json:"nonceStr"`
	TimeStamp    int64  `json:"timeStamp"`
	PackageValue string `json:"packageValue"`
	Sign         string `json:"sign"`
}

type AlipayCredential struct {
	AlipayOrderInfo string `json:"orderInfo"`
}

type UpmpCredential struct {
	Tn   string `json:"tn"`
	Mode string `json:"mode"`
}
