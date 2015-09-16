package main

import (
	"encoding/json"
	"fmt"
	"log"
	pingpp "pingpp-go/pingpp"
	"pingpp-go/pingpp/cardNumber"
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

func ExampleCardNumber_get() {
	param := &pingpp.CardQueryParams{
		App: "app_LibTW1n1SOq9Pin1",
	}

	card, err := cardNumber.Get("6222022003008481261", param)
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func main() {
	ExampleCardNumber_get()
}
