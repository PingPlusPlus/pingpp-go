/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package main

import (
	//"encoding/json"
	"fmt"
	"log"
	pingpp "pingpp-go/pingpp"
	"pingpp-go/pingpp/event"
)

func init() {
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	fmt.Println("Go SDK Version:", pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleEvent_get() {
	eve, err := event.Get("evt_zRFRk6ekazsH7t7yCqEeovhk")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eve)
}

func ExampleEvent_list() {
	params := &pingpp.EventListParams{}
	params.Filters.AddFilter("type", "", "charge.succeeded")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := event.List(params)
	for i.Next() {
		c := i.Event()
		fmt.Println(c)
	}
}

func main() {
	ExampleEvent_get()
}
