package main

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	pingpp "pingpp-go/pingpp"
	"pingpp-go/pingpp/charge"
	"strconv"
	"time"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_test_zL0abDjXX1mP4qLinL5y5mPG"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
	//设置商户的私钥 记得在Ping++上配置公钥
	pingpp.AccountPrivateKey = `
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
`
}

func ExampleCharge_new() {
	//针对metadata字段，可以在每一个 charge 对象中加入订单的一些详情，如颜色、型号等属性
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//metadata["type"] = "shoes"
	//metadata["size"] = "40"
	//extra 参数根据渠道不同有区别，下面注释的是一部分的示例
	extra := make(map[string]interface{})

	// upacp_wap
	// extra["result_url"] = "http://www.yourdomain.com"

	// bfb_wap
	// extra["result_url"] = "http://www.yourdomain.com"
	// extra["bfb_login"] = false

	// yeepay_wap
	// extra["product_category"] = "1"
	// extra["identity_id"] = "sadfadsjkfhasuidfhbjdasf"
	// extra["identity_type"] = 1
	// extra["terminal_type"] = 1
	// extra["terminal_id"] = "1sdf"
	// extra["user_ua"] = "1qwec"
	// extra["result_url"] = "http://www.yourdomain.com"

	// alipay_wap
	// extra["cancel_url"] = "http://www.yourdomain.com"
	// extra["success_url"] = "http://www.yourdomain.com"

	// wx_pub
	// extra["open_id"] = "sdafdgagfd"

	//jdpay_wap
	// extra["success_url"] = "http://www.yourdomain.com"
	// extra["fail_url"] = "http://www.yourdomain.com"
	// extra["token"] = "dsafadsfasdfadsjuyhfnhujkijunhaf"

	// wx_pub_qr
	// extra["product_id"] = "23sf"

	//  cnp_u或者cnp_f
	//	extra["source"] = "tok_fX1y1KifbPS0DqD4C8HOyH84"
	//	sms_code := make(map[string]interface{})
	//	extra["sms_code"] = sms_code
	//	sms_code["id"] = "sms_av900KPWrL84LqHqv5O44W54"
	//	sms_code["code"] = "123456"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	params := &pingpp.ChargeParams{
		Order_no:  strconv.Itoa(orderno),
		App:       pingpp.App{Id: "app_CyfHGK8eXPuL9uj9"},
		Amount:    1000,
		Channel:   "alipay_wap",
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body",
		Extra:     extra,
		Metadata:  metadata,
	}


	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	ch, err := charge.New(params)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(ch)

}

func ExampleCharge_get() {
	ch, err := charge.Get("ch_ejbLGCCaDWjT0ijzDSybL0mT")
	if err != nil {
		log.Fatal(err)
	}
	chstring, _ := json.Marshal(ch)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(chstring))
}

func ExampleCharge_list() {

	params := &pingpp.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := charge.List(params)
	for i.Next() {
		c := i.Charge()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func main() {
	ExampleCharge_new()
}
