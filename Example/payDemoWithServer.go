package main

//该示例只是模拟服务器端，可配合客户端体验完整的支付流程

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	// utils "pingpp/pingpp/utils"
	"strconv"
	"strings"
	"time"
)

var mymux *http.ServeMux

func Run() {

	//服务器路由
	mymux = http.NewServeMux()
	bind()

	err := http.ListenAndServe(":5623", mymux) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
	bind the router
*/
func bind() {
	mymux.HandleFunc("/pay", pay)
}

func pay(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		http.NotFound(w, r)
	} else if strings.ToUpper(r.Method) == "POST" {
		var chargeParams pingpp.ChargeParams
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		json.Unmarshal(buf.Bytes(), &chargeParams)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		orderno := r.Intn(999999999999999)
		extra := make(map[string]interface{})
		switch strings.ToLower(chargeParams.Channel) {
		case "upacp_wap":
			extra["result_url"] = "http://www.yourdomain.com/result"
		case "alipay_wap":
			extra["cancel_url"] = "http://www.yourdomain.com/cancel"
			extra["success_url"] = "http://www.yourdomain.com/success"
		case "bfb_wap":
			extra["result_url"] = "http://www.yourdomain.com/result"
			extra["bfb_login"] = false
		case "yeepay_wap":
			extra["product_category"] = "1"
			extra["identity_id"] = "your_identity_id"
			extra["identity_type"] = 1
			extra["terminal_type"] = 1
			extra["terminal_id"] = "1sdf"
			extra["user_ua"] = "1qwec"
			extra["result_url"] = "http://www.yourdomain.com/result"
		case "wx_pub":
			extra["open_id"] = "your_openid"
		case "jdpay_wap":
			extra["success_url"] = "http://www.yourdomain.com/success"
			extra["fail_url"] = "http://www.yourdomain.com/fail"
			extra["token"] = "your_token_from_jd"
		case "wx_pub_qr":
			extra["product_id"] = "your_productid"

		}

		pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"

		params := &pingpp.ChargeParams{
			Order_no:  strconv.Itoa(orderno),
			App:       pingpp.App{Id: "app_1Gqj58ynP0mHeX1q"},
			Amount:    chargeParams.Amount,
			Channel:   strings.ToLower(chargeParams.Channel),
			Currency:  "cny",
			Client_ip: "127.0.0.1",
			Subject:   "Your Subject",
			Body:      "Your Body",
			Extra:     extra,
		}

		//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
		ch, err := charge.New(params)

		if err != nil {
			errs, _ := json.Marshal(err)
			fmt.Fprint(w, string(errs))
		} else {
			chs, _ := json.Marshal(ch)
			fmt.Fprintln(w, string(chs))
		}

	}
}

func main() {
	Run()
}
