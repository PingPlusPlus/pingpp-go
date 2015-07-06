package main

import (
	"encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/redEnvelope"
	"log"
)

func init() {
	pingpp.LogLevel = 2
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	fmt.Println("Go SDK Version:", pingpp.Version())
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleRedEnvelope_new() {
	extra := make(map[string]interface{})
	extra["nick_name"] = "Nick Name"
	extra["send_name"] = "Send Name"
	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	redenvelopeParams := &pingpp.RedEnvelopeParams{
		App:         pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
		Channel:     "wx_pub",
		Order_no:    strconv.Itoa(orderno),
		Amount:      100,
		Currency:    "cny",
		Recipient:   "youropenid",
		Subject:     "Your Subject",
		Body:        "Your Body",
		Description: "Your Description",
		Extra:       extra,
	}
	redEnvelope, err := redEnvelope.New(redenvelopeParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(redEnvelope)
	redString, _ := json.Marshal(redEnvelope)
	fmt.Println(string(redString))
}

func ExampleRedEnvelope_get() {
	redEnvelope, err := redEnvelope.Get("red_id")
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
	ExampleRedEnvelope_new()
}
