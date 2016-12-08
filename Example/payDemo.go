/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package main

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	//"io/ioutil"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"

	// 直接读取获取直接配置
	//	privateKey, err := ioutil.ReadFile("your_rsa_private_key.pem")
	//	if err != nil {
	//		fmt.Errorf("read failure: %v", err)
	//	}
	//	pingpp.AccountPrivateKey = string(privateKey)

	//设置商户的私钥 记得在Ping++上配置公钥
	pingpp.AccountPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAx2MktxcKBEqdYRi2IgYcupPQIN5cxgiBL5udCCBJBNBbXPaq
uOE1qspfhB1KUzHXATnCONiSzubLcBTnwi2tz0ErRCeJZSERRCpbKx4eu6b1neUT
Wkga7xpZxWONEvkmZo5Nlhf4fXRPUYnO/bdGCNGpQ/HSJfWLtzmhCqO1aJwVhcDm
DMYz4bTkZavhFBdVyXf/8n7UKylk03eymlKJ1swQpeFcxaKfzsk1mJU7mc93mCWj
aR+VWkNbw4AQHDyHgbzH+zYARzCluiy5hXdixGEP+iO4ZBk48rEs1hKTvGz1k+jh
LCdkdpBRjq0pK/htjA3Ce8pF2AJs+fgN6ZUumQIDAQABAoIBAFa4MEfRpXGoYjrQ
3KZ/sg8UKvmgvQkEuetS60GViSym0pXkUuyGRyk5S8HSW3lDvBe0X10KFRAYIXNm
JEa4R1hVJ9REveVWNIRJR83BE+zZ+QnrkDc8FTrZYyIO4lTWOHVyfxxA4Lrv02/L
WFPRWoyLY+tBSf1ohpPyZLCT81rDglT1Z4svX020y8tXvnQqQiOjl4q7Zu4b26HU
TQ463ntMEhM5u7y9MFcxGRaOpF/gARlMGqDu6T8h/oYMiOSLoXOuTR7B80yaX/Mj
RZfUBoZMb5thX9qBLQ7dYnTkwaxwerYPrYvQrW9vtsswZ5NeIbEmCZyorUe8DOmQ
hT1+HmECgYEA/iQERHhZKHXnP0gvhl/uEOGOvLjD5H1D6zClzOHMmOcIF5OuEQb0
VcSMV+8emN7SCp/b/LVgKa27Mla9eXm+EXABRFcI7qGYsYXfbCD7EYX3TaJSp/30
jyLBy+MsHCTEiLeylSh7kHqgTR8tKND8UIzXo9aM7JqwFqleeXGyh7MCgYEAyNiU
EUzyBAv9sui3ZgVYRiVvTilk2HVTY6u61/mMOLsTrX3eYQaqb4GRJJShJO9mmsxX
RHBEZQJvUqqF9PapOsyv8HKuF5+UP6svHnJo7sn9gCvV/h1HTHqzFcYSvUaXnrym
D/0Tthf8CDeuGp5UFWMoFZF14HTr1oQROGAASoMCgYA0bZmzxmAeSLR8CZhEUGX8
dYvMwxEmgfERA+gwbCSZJpA0zPKL8LNXPkT1nw7g2pbaOkBX0dMUxhJoQBy2grcD
QegBATOGhy/I76U32VXyN4DdMy96GJnrLXBtb2AaLjudOMhOnRtgouuO/W+DjBmB
RIz377sC1KafBjHHO/1ooQKBgDQqfJrZv2ppquVTKH9pF/pwMq68daL7JkOXERqT
iGYbwQqozJ+q2Y3Iu2gi6o/rVl0SggAWoM0TitKP0+dCQcYx7+imAK3GFv1KexyP
Xs3WzO8Dc7ti42fr3qPjJG7g7PSfzwoME5iSNjX0MFZdlT1Q2dJwS4uXEsJO3yIj
XS/9AoGBALRApgtUA7Odw4tjCLGvxXuLFnyRkg6hFqoXAP2j8H9bJDOlSSVwQTFd
ahbcIDtQJS57vXUGK2uspbFKLm1WCFzPVyuxDIW6oue/kO+YxxU3NA58zk8oaORq
eA3YvHc7ZmRjVnVkxnXjKofrL6jF5A+lXSXnXchrv2ZYI+1pOsIV
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
		App:       pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
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
	ch, err := charge.Get("ch_a9CmfHTGGaz1urHiL8m5OiX1")
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
	//app[id]是必须参数
	params.Filters.AddFilter("app[id]", "", "app_1Gqj58ynP0mHeX1q")
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
	//ExampleCharge_get()
	ExampleCharge_list()
}
