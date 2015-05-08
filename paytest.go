package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pingplusplus/pingpp-go/pingpp/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/pingpp/charge"
	"log"
	"math/rand"
	"net/http"
	// utils "pingpp/pingpp/utils"
	"strconv"
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
	if r.Method == "GET" {
		http.NotFound(w, r)
	} else if r.Method == "POST" {
		var chargeParams pingpp.ChargeParams
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &chargeParams)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		orderno := r.Intn(999999999999999)
		pingpp.Key = "YOUR-KEY"

		// client := &http.Client{}
		// wx_app_id := "wx95df8698f12122c7"
		// redirect_url := "http://test.pinpula.com/wx-pub/test_sdk.php?showwxpaytitle=1"
		// codeUrl := utils.CreateOauthUrlForCode(wx_app_id, redirect_url, false)
		// fmt.Fprintln(w, codeUrl)
		// req, _ := http.NewRequest("GET", "http://test.pinpula.com/wx-pub/test_go.php", nil)
		// urlStr := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx95df8698f12122c7&redirect_url=http%3A%2F%2Ftest.pinpula.com%2Fwx-pub%2Ftest_sdk.php%3Fshowwxpaytitle%3D1&response_type=code&scope=snsapi_base&state=STATE%23wechat_redirect"
		// req.Header.Add("location", urlStr)
		// resp, _ := client.Do(req)
		// fmt.Println(req.URL)
		// fmt.Println(resp)
		params := &pingpp.ChargeParams{
			Order_no:  strconv.Itoa(orderno),
			App:       pingpp.App{Id: "YOUR-APP-ID"},
			Amount:    chargeParams.Amount,
			Channel:   chargeParams.Channel,
			Currency:  "cny",
			Client_ip: "127.0.0.1",
			Subject:   "Your Subject",
			Body:      "Your Body",
			Extra:     pingpp.Extra{Open_id: "your open_id"},
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

		//获取 jssdk所需的签名
		// wx_app_secret := "6c521c5a75c135e094b7d90dc5600b7a"
		// jsTicket := utils.GetJsapiTicket(wx_app_id, wx_app_secret)
		// urls := "http://test.pinpula.com/wx-pub/test_go.php"
		// signatures := utils.GetSignature(ch, jsTicket, urls)
		// maps := make(map[string]interface{})
		// maps["charge"] = ch
		// maps["signature"] = fmt.Sprintf("%x", signatures)
		// a, _ := json.Marshal(maps)
		// fmt.Fprintf(w, string(a))
	}
}

func main() {
	Run()
}
