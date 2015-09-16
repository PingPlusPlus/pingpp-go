package main

import (
	"encoding/json"
	"fmt"
	"log"
	pingpp "pingpp-go/pingpp"
	"pingpp-go/pingpp/card"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_live_vjfr90jj1q985KuPO84iP8KO"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleCard_new() {

	cus_id := "cus_9K4KS8jLKq50yP"

	// params := make(map[string]interface{})
	// params["card_number"] = "6222022003008481261"
	// params["brand"] = "UnionPay"
	// params["funding"] = "debit"
	// params["bank"] = "icbc"
	// params["name"] = "张三"
	// params["cred_type"] = "ID"
	// params["cred_number"] = "350583199009153732"
	// params["phone_number"] = "13045678901"

	param := &pingpp.CardParams{
		Source: "tok_u5qrjPjL4yH4zz9WPKjjjvvT",
	}

	card, err := card.New(cus_id, param)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(card)

}

func ExampleCard_get() {
	card, err := card.Get("cus_9K4KS8jLKq50yP", "card_brvbP0rLuPW5G440W5jXjjT4")
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func ExampleCard_list() {

	params := &pingpp.CardListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	cus_id := "cus_qn5enTSyzPOGfz"
	i := card.List(cus_id, params)
	for i.Next() {
		c := i.Card()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func ExampleCard_delete() {
	card, err := card.Delete("cus_9K4KS8jLKq50yP", "card_brvbP0rLuPW5G440W5jXjjT4")
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func main() {
	ExampleCard_list()
}
