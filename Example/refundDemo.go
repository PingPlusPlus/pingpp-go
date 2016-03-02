package main

import (
	"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/refund"
	"log"
)

func init() {
	pingpp.LogLevel = 2
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	fmt.Println("Go SDK Version:", pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
	//设置商户的私钥 记得在Ping++上配置公钥
	//pingpp.AccountPrivateKey
}

func ExampleRefund_new() {
	params := &pingpp.RefundParams{
		Amount:      1, //可以注释不上传
		Description: "12345",
	}
	re, err := refund.New("ch_id", params) //ch_id 是已付款的订单号

	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(re)
	log.Printf("%v\n", string(restring))
}

func ExampleRefund_get() {
	re, err := refund.Get("ch_id", "refund_id")
	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(re)

	log.Printf("%v\n", string(restring))
}

func ExampleRefund_list() {
	params := &pingpp.RefundListParams{}
	// params.Filters.AddFilter("limit", "", "1")
	// //设置是不是只需要之前设置的 limit 这一个查询参数
	// params.Single = true
	i := refund.List("ch_id", params)
	for i.Next() {
		c := i.Refund()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	fmt.Println(i)
}

func main() {
	ExampleRefund_new()
}
