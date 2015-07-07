package main

import (
	"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	"log"
	"math/rand"
	"strconv"
	"time"
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
}

func ExampleCharge_new() {
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//extra 参数根据渠道不同有区别，下面注释的是一部分的示例
	extra := make(map[string]interface{})
	// //upacp_wap
	// extra["result_url"] = "http://www.yourdomain.com"

	//bfb_wap
	// extra["result_url"] = "http://www.yourdomain.com"
	// extra["bfb_login"] = false

	// //yeepay_wap
	// extra["product_category"] = "1"
	// extra["identity_id"] = "sadfadsjkfhasuidfhbjdasf"
	// extra["identity_type"] = 1
	// extra["terminal_type"] = 1
	// extra["terminal_id"] = "1sdf"
	// extra["user_ua"] = "1qwec"
	// extra["result_url"] = "http://www.yourdomain.com"

	// //alipay_wap
	// extra["cancel_url"] = "http://www.yourdomain.com"
	// extra["success_url"] = "http://www.yourdomain.com"

	// //wx_pub
	// extra["open_id"] = "sdafdgagfd"

	//jdpay_wap
	// extra["success_url"] = "http://www.yourdomain.com"
	// extra["fail_url"] = "http://www.yourdomain.com"
	// extra["token"] = "dsafadsfasdfadsjuyhfnhujkijunhaf"

	// //wx_pub_qr
	// extra["product_id"] = "23sf"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	params := &pingpp.ChargeParams{
		Order_no:  strconv.Itoa(orderno),
		App:       pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
		Amount:    1000,
		Channel:   "alipay",
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
	ch, err := charge.Get("ch_ejbLGCCaDWjT0ijzDSybL0mT")
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
	ExampleCharge_new()
}
