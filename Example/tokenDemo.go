package main

import (
	"encoding/json"
	"fmt"
	"log"
	pingpp "pingpp-go/pingpp"
	"pingpp-go/pingpp/token"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 4
	//设置 API Key
	pingpp.Key = "sk_live_vjfr90jj1q985KuPO84iP8KO"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleToken_new() {

	// cus_id := "cus_9K4KS8jLKq50yP"

	params := make(map[string]interface{})
	params["card_number"] = "6222022003008481261"
	params["brand"] = "UnionPay"
	params["funding"] = "debit"
	params["bank"] = "icbc"
	params["name"] = "张三"
	params["cred_type"] = "ID"
	params["cred_number"] = "350583199009153732"
	params["phone_number"] = "13045678901"

	param := &pingpp.TokenParams{
		Card:     params,
		Order_no: "yyyuuuiiii",
		Amount:   1,
		App:      "app_LibTW1n1SOq9Pin1",
	}

	card, err := token.New(param)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(card)

}

func ExampleToken_get() {
	tok, err := token.Get("tok_TenD0KTerf50PerXfHbDOG0K")
	if err != nil {
		log.Fatal(err)
	}
	tokstring, _ := json.Marshal(tok)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(tokstring))
}

func main() {
	ExampleToken_new()
}
