//本示例使用 go 直接调用 api charge 的 create 接口，而不是使用 go-sdk。
//返回的也跟sdk中返回的是对象不一样.这里直接返回了 json 字符串。
// 其他接口调用方法跟本示例类似，具体详见[API 文档](https://pingxx.com/document/api)
// 建议使用 go-sdk,会比较方便一些。

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Charge struct {
	Id              string                 `json:"id"`
	Object          string                 `json:"object"`
	Created         uint64                 `json:"created"`
	Livemode        bool                   `json:"livemode"`
	Paid            bool                   `json:"paid"`
	Refunded        bool                   `json:"refunded"`
	App             string                 `json:"app"`
	Channel         string                 `json:"channel"`
	Order_no        string                 `json:"order_no"`
	Client_ip       string                 `json:"client_ip"`
	Amount          int                    `json:"amount"`
	Amount_settle   int64                  `json:"amount_settle"`
	Currency        string                 `json:"currency"`
	Subject         string                 `json:"subject"`
	Body            string                 `json:"body"`
	Extra_data      Extra                  `json:"extra"`
	Time_paid       uint64                 `json:"time_paid"`
	Time_expire     uint64                 `json:"time_expire"`
	Time_settle     uint64                 `json:"time_settle"`
	Transaction_no  string                 `json:"transaction_no"`
	Refunds         RefundList             `json:"refunds"`
	Amount_refunded int64                  `json:"amount_refunded"`
	Failure_code    string                 `json:"failure_code"`
	Failure_msg     string                 `json:"failure_msg"`
	Metadata        map[string]interface{} `json:"metadata"`
	Credential      map[string]interface{} `json:"credential"`
	Description     string                 `json:"description"`
}

type RefundList struct {
	Object   string   `json:"object"`
	Url      string   `json:"url"`
	Has_more bool     `json:"has_more"`
	Refunds  []Refund `json:"data"`
}

type Refund struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Order_no     string            `json:"order_no"`
	Amount       int64             `json:"amount"`
	Succeed      bool              `json:"succeed"`
	Created      uint64            `json:"created"`
	Time_succeed uint64            `json:"time_succeeded"`
	Description  string            `json:"description"`
	Failure_code int64             `json:"failure_code"`
	Failure_msg  string            `json:"failure_message"`
	Metadata     map[string]string `json:"metadata"`
	Charge_id    string            `json:"charge"`
}

type ChargeParams struct {
	OrderNo    string `json:"order_no"`
	Amount     uint64 `json:"amount"`
	Channel    string `json:"channel"`
	Currency   string `json:"currency"`
	App        *App   `json:"app"`
	Client_ip  string `json:"client_ip"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Extra_data *Extra `json:"extra,omitempty"`
	Descripton string `json:"description,omitempty"`
}

type App struct {
	Id string `json:"id"`
}

type Extra struct {
	Result_url    string `json:"result_url,omitempty"`
	Success_url   string `json:"success_url,omitempty"`
	Cancel_url    string `json:"cancel_url,omitempty"`
	Trade_type    string `json:"trade_type,omitempty"`
	Open_id       string `json:"openid,omitempty"`
	Bfb_login     bool   `json:"bfb_login,omitempty"`
	Payment_token string `json:"payment_token,omitempty"`
}

func chargeNew() {
	createParams := &ChargeParams{
		OrderNo:   "dsafsdk435",
		Amount:    100,
		Channel:   "wx",
		Currency:  "cny",
		App:       &App{Id: "YOUR-APP-ID"},
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body"}

	chargeParams, _ := json.Marshal(createParams)
	body := bytes.NewBuffer([]byte(chargeParams))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.pingxx.com/v1/charges", body)
	req.SetBasicAuth("YOUR-KEY", "")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(result))
	//返回给客户端的一定要是charge 的 json 字符串。
	//下面几行代码是把json转换成struct的示例，方便读取其中的字段
	var charge Charge
	json.Unmarshal(result, &charge)
	fmt.Println(&charge)
}

func main() {
	chargeNew()
}
