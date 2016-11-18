package main

import (
	"encoding/json"
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/token"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 4
	//设置 API Key
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleToken_get() {
	tok, err := token.Get("tok_ALeVGaNi5marUpMu4bapUZSZ")
	if err != nil {
		log.Fatal(err)
	}
	tokstring, _ := json.Marshal(tok)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(tokstring))
}

func main() {
	ExampleToken_get()
}
