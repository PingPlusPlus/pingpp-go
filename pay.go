package main

import (
	"encoding/json"
	"fmt"
	"log"
	"pingpp/pingpp"
	"pingpp/pingpp/charge"
	"pingpp/pingpp/redEnvelope"
	"pingpp/pingpp/refund"
)

func init() {
	pingpp.Key = "YOUR-KEY"
}

func ExampleCharge_new() {
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	params := &pingpp.ChargeParams{
		Order_no:  "sdafewdhtydr",
		App:       pingpp.App{Id: "YOUR-APP-ID"},
		Amount:    100,
		Channel:   "wx",
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body",
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
	chs, err := json.Marshal(ch)
	if err != nil {
		fmt.Println(err)
	} else {
		//打印 charge json
		fmt.Println(string(chs))
	}
}

func ExampleCharge_get() {
	ch, err := charge.Get("ch_id")
	if err != nil {
		log.Fatal(err)
	}
	chstring, _ := json.Marshal(ch)
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

func ExampleRefund_new() {
	params := &pingpp.RefundParams{
		Amount:      1,
		Description: "12345",
	}
	re, err := refund.New("ch_id", params)

	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(re)
	log.Printf("%v\n", string(restring))
}

func ExampleRefund_get() {
	re, err := refund.Get("ch_id", "re_id")
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

func ExampleRedEnvelope_new() {
	redenvelopeParams := &pingpp.RedEnvelopeParams{
		App:         &pingpp.App{Id: "YOUR-APP-ID"},
		Channel:     "wx_pub",
		Order_no:    "987654345",
		Amount:      100,
		Currency:    "cny",
		Recipient:   "o9zpMs9jIaLynQY9N6yYcmxcZ2zc",
		Subject:     "Your Subject",
		Body:        "Your Body",
		Description: "Your Description",
		Extra:       pingpp.RedEnvelopeExtra{Nick_name: "Nick Name", Send_name: "Send Name"},
	}
	redEnvelope, err := redEnvelope.New(redenvelopeParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(redEnvelope)
}

func ExampleRedEnvelope_get() {
	redEnvelope, err := redEnvelope.Get("RED_ID")
	if err != nil {
		log.Fatal(err)
	}
	restring, _ := json.Marshal(redEnvelope)
	log.Printf("%v\n", string(restring))
}

func ExampleRedEnvelope_list() {
	params := &pingpp.RedEnvelopeListParams{}
	params.Filters.AddFilter("limit", "", "2")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := redEnvelope.List(params)
	for i.Next() {
		c := i.RedEnvelope()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
}

func main() {
	ExampleCharge_new()
}
