package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	pingpp "pingpp-go/pingpp"
)

// 商户的私钥，还有一个公钥需要上传至Ping++
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEoQIBAAKCAQEAgwaTUmUK9BxxSFjDNvD+zOlJlGZTgucF6aFQWmtc7UBGdRpr
fH7Bs+/Gvf2wAbNJAcwLF3DVEl4uP4rqatRkCQLLUwsag4haJ/QhACP9XCxNqvBh
3Z6HowLgQvQ7PXGiIVsrqs11HhEzpq01aMxuBuDnujZ/+j1cO2ibnsbL99hXAPW0
LRicUEvHzPtduRvzKBaeHosnVaHYS8l7Yf9ExdmP7dA9vOqnDjJ0nxzQ6NAIbyi4
kDgUjn14HFqUnkCVqmtyMr6Haq1BZx0oH0V35dW3Ew1hu1BLAvZcNJB7LQHOLeKF
vmUA9Euwypho71ZH2NYDATg07JK4g3Tj2k0ggwIBJQKCAQBxUcuFh9IfOzGKraHq
U+NJb9DvDGPm1aQ/rhv7LGwLddwSQFz2DMMz1ksFLm61XMmFA3823iAP4sA29K8Q
SPySK+4eTsrSn3CRQddaEUnTLTxAz+XwGmd4Oda2cly/hN8/cXHf1Emddqk2NO/r
7xMM3jAyWKYI2x9V+ZtY40ie/wXpBHjuSOJvJDcif8x+moZ9OQpg+fvf9H0Oekxv
MDs4QEs6cHawGYUZr7R5S1XrgWLTWPfaEqvZ1mSNcavoB/sCaixBc5vBlf8SUpfS
IkuFo0IUQRwjzt2ibU1BA1F0V4RP2DGNkU9bVV2bwgPkz0/1hQagvfmtx2WpJxzy
/3WtAoGBAObnPFc168sSWT4EJhZcxpbLl3yY06AsxaWgFcbUMvafFehlHsz0atzK
smG+oqSXh1Gs0ENsZcs1y8egI7h+/aZlF+eHJEbXcZHq7q8geUzd/kskdThHc/3f
EEwJX5OwtBbxv8xJuzgESuR1ucj6GVZW2WLR86eON27wLZzDZHILAoGBAJFETDEj
qIttcKoI8YLp8rDbtS2Flnvt6YIS9kwtPImtuVCNH2eM4khdzOo0OxD4+xx9+ZaP
rLKk0cw4q21a+bUKiMd/7JR7ZkCO/vlz1IO1UqeHLODqdRzgF67YBRFNyIY/0MAo
+cUuZhgYxkt0rT0OGtCFaOStbLwHsg2d8U5pAoGBAKh/Od7Ge6IGeHlc+TLpwVlI
dXaZCSHbe3jcmkT7tnydml16ORIaJHea1TKSB/uRK2UdPgev90FC7qZt8I2NGfX2
vmquuZt6nvvHHOCo/pIJx2dEHi/8ykOixrQG1wrousSikuEu4pCpL7t4gKdqZYQx
ieBUBNQ3WOJAj/zF71ojAoGARquasB8vZmyXrKpnp3i7QUhKTX9CSiC2yahcJQ8W
iCsANQZiTg0v0C2h9WWEhMwnIp4mZOvzJnmyK//ksb2OPGX9WiKOxMZpGHzqvotg
d2z3xx8qlu6Z1rI1CPN4FkGEJaKBOuOAbcOLnQUiMovJ5lnjiAmM+aB6Fkju1jEp
SL0CgYA/NVZe5W7VbC0HxHGA8KyuSiqY3uWrXiE+jHMNFNtSi5LhR0a9gaxrxa58
hTh6REbOCFcQjX9OCx65RG+wn6X0katVFu+x1R+kZ7A1dGbOthXisocLHN58Cjf3
EsWDR34yF7Dr4h0NyIzQ6UmeKHZoXo+tJ0wOq8gl5uN52uCXSQ==
-----END RSA PRIVATE KEY-----
`)

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
	//签名是 header的 x-pingplusplus-signature中的值
	signed := `PcU0SMJhbPObiIVinNnalZOjI02koWozxLrxa3WQW3rK/n7I+EuVGuXvhsq2MIfUaNiHZDgRFYybGtKr1uuFzEXjA4PwmnDHfWgwRPdjgseoU0eke6ZqGpklBRVTbF6PUy6/vAqur4xb7h1wpdrteUpCPafzDmVPsQLicdojJ/TF9ACjQW8gTNiS6tE9gL5hxy0RJ3/okRJo6dz2pvJBWkjCrgp/r98z/LQijA1o//atZrH63+DcL/GwEOgaymqbodzusXF+g6WMJ/GTJgjdPRHvpO9UAAUKkOQqvwthJvsXIH/L1xqvy+tFpo2J0Ptwg85bowKoyy1qC5ak3sqWqw==`

	//待验签的数据
	data := `{"id":"evt_04qN8cXQvIhssduhS4hpqd9p","created":1427555016,"livemode":false,"type":"account.summary.available","data":{"object":{"acct_id":"acct_0eHSiDyzv9G09ejT","object":"account_daily_summary","acct_display_name":"xx公司","created":1425139260,"summary_from":1425052800,"summary_to":1425139199,"charges_amount":1000,"charges_count":100}},"object":"event","pending_webhooks":2,"request":null,"scope":"acct_1234567890123456","acct_id":"acct_1234567890123456"}`

	//Ping++ 公钥
	publicKey, err := ioutil.ReadFile("my-server.pub")
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
func genRsaSign() {
	data := []byte(`www.pingxx.com`)
	sign, err := pingpp.GenSign(data, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	encodeSign := base64.StdEncoding.EncodeToString(sign)
	fmt.Println(encodeSign)
}
