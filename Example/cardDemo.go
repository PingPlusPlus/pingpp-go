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
	pingpp.LogLevel = 4
	//设置 API Key
	pingpp.Key = "sk_test_zL0abDjXX1mP4qLinL5y5mPG"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleCard_new() {

	cus_id := "cus_ALeWGZ8lsN9Czk"

	param := &pingpp.CardParams{
		Source: "tok_AMBKETCThoW7nYUBgpnvhwfu",
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
	card, err := card.Get("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
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
	cus_id := "cus_ALeWGZ8lsN9Czk"
	i := card.List(cus_id, params)
	for i.Next() {
		c := i.Card()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func ExampleCard_delete() {
	card, err := card.Delete("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func main() {
	ExampleCard_new()
}
