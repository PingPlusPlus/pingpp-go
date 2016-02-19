/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Go SDK 使用，只是提供一个参考。
 */
package main

import (
	"encoding/base64"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"io/ioutil"
)

// 商户的私钥(可以放在变量中或保存到文件中)，还有一个公钥需要上传至Ping++
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIIEoQIBAAKCAQEAgwaTUmUK9BxxSFjDNvD+zOlJlGZTgucF6aFQWmtc7UBGdRpr
//fH7Bs+/Gvf2wAbNJAcwLF3DVEl4uP4rqatRkCQLLUwsag4haJ/QhACP9XCxNqvBh
//3Z6HowLgQvQ7PXGiIVsrqs11HhEzpq01aMxuBuDnujZ/+j1cO2ibnsbL99hXAPW0
//LRicUEvHzPtduRvzKBaeHosnVaHYS8l7Yf9ExdmP7dA9vOqnDjJ0nxzQ6NAIbyi4
//kDgUjn14HFqUnkCVqmtyMr6Haq1BZx0oH0V35dW3Ew1hu1BLAvZcNJB7LQHOLeKF
//vmUA9Euwypho71ZH2NYDATg07JK4g3Tj2k0ggwIBJQKCAQBxUcuFh9IfOzGKraHq
//U+NJb9DvDGPm1aQ/rhv7LGwLddwSQFz2DMMz1ksFLm61XMmFA3823iAP4sA29K8Q
//SPySK+4eTsrSn3CRQddaEUnTLTxAz+XwGmd4Oda2cly/hN8/cXHf1Emddqk2NO/r
//7xMM3jAyWKYI2x9V+ZtY40ie/wXpBHjuSOJvJDcif8x+moZ9OQpg+fvf9H0Oekxv
//MDs4QEs6cHawGYUZr7R5S1XrgWLTWPfaEqvZ1mSNcavoB/sCaixBc5vBlf8SUpfS
//IkuFo0IUQRwjzt2ibU1BA1F0V4RP2DGNkU9bVV2bwgPkz0/1hQagvfmtx2WpJxzy
///3WtAoGBAObnPFc168sSWT4EJhZcxpbLl3yY06AsxaWgFcbUMvafFehlHsz0atzK
//smG+oqSXh1Gs0ENsZcs1y8egI7h+/aZlF+eHJEbXcZHq7q8geUzd/kskdThHc/3f
//EEwJX5OwtBbxv8xJuzgESuR1ucj6GVZW2WLR86eON27wLZzDZHILAoGBAJFETDEj
//qIttcKoI8YLp8rDbtS2Flnvt6YIS9kwtPImtuVCNH2eM4khdzOo0OxD4+xx9+ZaP
//rLKk0cw4q21a+bUKiMd/7JR7ZkCO/vlz1IO1UqeHLODqdRzgF67YBRFNyIY/0MAo
//+cUuZhgYxkt0rT0OGtCFaOStbLwHsg2d8U5pAoGBAKh/Od7Ge6IGeHlc+TLpwVlI
//dXaZCSHbe3jcmkT7tnydml16ORIaJHea1TKSB/uRK2UdPgev90FC7qZt8I2NGfX2
//vmquuZt6nvvHHOCo/pIJx2dEHi/8ykOixrQG1wrousSikuEu4pCpL7t4gKdqZYQx
//ieBUBNQ3WOJAj/zF71ojAoGARquasB8vZmyXrKpnp3i7QUhKTX9CSiC2yahcJQ8W
//iCsANQZiTg0v0C2h9WWEhMwnIp4mZOvzJnmyK//ksb2OPGX9WiKOxMZpGHzqvotg
//d2z3xx8qlu6Z1rI1CPN4FkGEJaKBOuOAbcOLnQUiMovJ5lnjiAmM+aB6Fkju1jEp
//SL0CgYA/NVZe5W7VbC0HxHGA8KyuSiqY3uWrXiE+jHMNFNtSi5LhR0a9gaxrxa58
//hTh6REbOCFcQjX9OCx65RG+wn6X0katVFu+x1R+kZ7A1dGbOthXisocLHN58Cjf3
//EsWDR34yF7Dr4h0NyIzQ6UmeKHZoXo+tJ0wOq8gl5uN52uCXSQ==
//-----END RSA PRIVATE KEY-----
//`)

//商户的公钥需要上传至Ping++
// var publicKey = []byte(`
// -----BEGIN PUBLIC KEY-----
// MIIBIDANBgkqhkiG9w0BAQEFAAOCAQ0AMIIBCAKCAQEAgwaTUmUK9BxxSFjDNvD+
// zOlJlGZTgucF6aFQWmtc7UBGdRprfH7Bs+/Gvf2wAbNJAcwLF3DVEl4uP4rqatRk
// CQLLUwsag4haJ/QhACP9XCxNqvBh3Z6HowLgQvQ7PXGiIVsrqs11HhEzpq01aMxu
// BuDnujZ/+j1cO2ibnsbL99hXAPW0LRicUEvHzPtduRvzKBaeHosnVaHYS8l7Yf9E
// xdmP7dA9vOqnDjJ0nxzQ6NAIbyi4kDgUjn14HFqUnkCVqmtyMr6Haq1BZx0oH0V3
// 5dW3Ew1hu1BLAvZcNJB7LQHOLeKFvmUA9Euwypho71ZH2NYDATg07JK4g3Tj2k0g
// gwIBJQ==
// -----END PUBLIC KEY-----
//`)

//验证 Webhook 签名
func main() {
	verifyRsaSign()
}

//目前只在webhook中使用
func verifyRsaSign() {
	//签名在头部信息的 x-pingplusplus-signature 字段
	signed := `BX5sToHUzPSJvAfXqhtJicsuPjt3yvq804PguzLnMruCSvZ4C7xYS4trdg1blJPh26eeK/P2QfCCHpWKedsRS3bPKkjAvugnMKs+3Zs1k+PshAiZsET4sWPGNnf1E89Kh7/2XMa1mgbXtHt7zPNC4kamTqUL/QmEVI8LJNq7C9P3LR03kK2szJDhPzkWPgRyY2YpD2eq1aCJm0bkX9mBWTZdSYFhKt3vuM1Qjp5PWXk0tN5h9dNFqpisihK7XboB81poER2SmnZ8PIslzWu2iULM7VWxmEDA70JKBJFweqLCFBHRszA8Nt3AXF0z5qe61oH1oSUmtPwNhdQQ2G5X3g==`

	//待验签的数据
	data := `{"id":"evt_eYa58Wd44Glerl8AgfYfd1sL","created":1434368075,"livemode":true,"type":"charge.succeeded","data":{"object":{"id":"ch_bq9IHKnn6GnLzsS0swOujr4x","object":"charge","created":1434368069,"livemode":true,"paid":true,"refunded":false,"app":"app_vcPcqDeS88ixrPlu","channel":"wx","order_no":"2015d019f7cf6c0d","client_ip":"140.227.22.72","amount":100,"amount_settle":0,"currency":"cny","subject":"An Apple","body":"A Big Red Apple","extra":{},"time_paid":1434368074,"time_expire":1434455469,"time_settle":null,"transaction_no":"1014400031201506150354653857","refunds":{"object":"list","url":"/v1/charges/ch_bq9IHKnn6GnLzsS0swOujr4x/refunds","has_more":false,"data":[]},"amount_refunded":0,"failure_code":null,"failure_msg":null,"metadata":{},"credential":{},"description":null}},"object":"event","pending_webhooks":0,"request":"iar_Xc2SGjrbdmT0eeKWeCsvLhbL"}`

	// 请从 https://dashboard.pingxx.com 获取「Ping++ 公钥」
	publicKey, err := ioutil.ReadFile("pingpp_rsa_public_key.pem")
	if err != nil {
		fmt.Errorf("read failure: %v", err)
	}

	//base64解码再验证
	decodeStr, _ := base64.StdEncoding.DecodeString(signed)
	errs := pingpp.Verify([]byte(data), publicKey, decodeStr)
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("success")
	}
}

//利用商户的私钥 生成RSA签名(SHA256)。 用于请求 Ping++ 的URL。Ping++接受后用商户上传给Ping++的公钥来验签
//func genRsaSign() {
//	data := []byte(`www.pingxx.com`)
//	sign, err := pingpp.GenSign(data, privateKey)
//	if err != nil {
//		fmt.Println(err)
//	}
//	encodeSign := base64.StdEncoding.EncodeToString(sign)
//	fmt.Println(encodeSign)
//}
