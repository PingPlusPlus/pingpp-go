/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/transfer"
)

func init() {
	pingpp.LogLevel = 2
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"

	fmt.Println("Go SDK Version:", pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
	//设置商户的私钥 记得在Ping++上配置公钥
	//pingpp.AccountPrivateKey
}

func ExampleTransfer_new() {
	extra := make(map[string]interface{})
	extra["user_name"] = "user name"
	extra["force_check"] = false

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成订单号，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	transferParams := &pingpp.TransferParams{
		App:         pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
		Channel:     "wx_pub",
		Order_no:    strconv.Itoa(orderno),
		Amount:      100,
		Currency:    "cny",
		Type:        "b2c",
		Recipient:   "youropenid",
		Description: "Your Description",
		Extra:       extra,
	}
	transfer, err := transfer.New(transferParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transfer)
	fr, _ := json.Marshal(transfer)
	fmt.Println(string(fr))
}

func ExampleTransfer_get() {
	transfer, err := transfer.Get("tr_98WLGObXbPKO1avLe5CqH0WH")
	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(transfer)
	log.Printf("%v\n", string(restring))
}

func ExampleTransfer_list() {
	params := &pingpp.TransferListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := transfer.List(params)
	for i.Next() {
		c := i.Transfer()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
}

func main() {
	ExampleTransfer_get()
}
