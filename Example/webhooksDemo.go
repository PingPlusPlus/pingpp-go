/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Go SDK 使用，只是提供一个参考。
 */
package main

import (
	"bytes"
// "encoding/json"
	"encoding/base64"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	pingpp "pingpp-go/pingpp"
	"strings"
)

func main() {
	var myserver *http.ServeMux
	myserver = http.NewServeMux()
	myserver.HandleFunc("/webhook", webhook)
	//设置监听的端口,5623 端口只是示例使用，具体端口号根据你的代码自行定义
	err := http.ListenAndServe(":5623", myserver)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		//示例 - 签名在头部信息的 x-pingplusplus-signature 字段
		//signed := `BX5sToHUzPSJvAfXqhtJicsuPjt3yvq804PguzLnMruCSvZ4C7xYS4trdg1blJPh26eeK/P2QfCCHpWKedsRS3bPKkjAvugnMKs+3Zs1k+PshAiZsET4sWPGNnf1E89Kh7/2XMa1mgbXtHt7zPNC4kamTqUL/QmEVI8LJNq7C9P3LR03kK2szJDhPzkWPgRyY2YpD2eq1aCJm0bkX9mBWTZdSYFhKt3vuM1Qjp5PWXk0tN5h9dNFqpisihK7XboB81poER2SmnZ8PIslzWu2iULM7VWxmEDA70JKBJFweqLCFBHRszA8Nt3AXF0z5qe61oH1oSUmtPwNhdQQ2G5X3g==`
		signed := r.Header.Get("X-Pingplusplus-Signature")

		//示例 - 待验签的数据
		//data := `{"id":"evt_eYa58Wd44Glerl8AgfYfd1sL","created":1434368075,"livemode":true,"type":"charge.succeeded","data":{"object":{"id":"ch_bq9IHKnn6GnLzsS0swOujr4x","object":"charge","created":1434368069,"livemode":true,"paid":true,"refunded":false,"app":"app_vcPcqDeS88ixrPlu","channel":"wx","order_no":"2015d019f7cf6c0d","client_ip":"140.227.22.72","amount":100,"amount_settle":0,"currency":"cny","subject":"An Apple","body":"A Big Red Apple","extra":{},"time_paid":1434368074,"time_expire":1434455469,"time_settle":null,"transaction_no":"1014400031201506150354653857","refunds":{"object":"list","url":"/v1/charges/ch_bq9IHKnn6GnLzsS0swOujr4x/refunds","has_more":false,"data":[]},"amount_refunded":0,"failure_code":null,"failure_msg":null,"metadata":{},"credential":{},"description":null}},"object":"event","pending_webhooks":0,"request":"iar_Xc2SGjrbdmT0eeKWeCsvLhbL"}`
		data := buf.String()
		// 请从 https://dashboard.pingxx.com 获取「Ping++ 公钥」
		publicKey, err := ioutil.ReadFile("pingpp_rsa_public_key.pem")
		if err != nil {
			fmt.Errorf("read failure: %v", err)
		}

		//base64解码再验证
		decodeStr, _ := base64.StdEncoding.DecodeString(signed)
		errs := pingpp.Verify([]byte(data), publicKey, decodeStr)
		if errs != nil {
			fmt.Println(errs)
			return
		} else {
			fmt.Println("success")
		}

		webhook, err := pingpp.ParseWebhooks(buf.Bytes())
		//fmt.Println(webhook.Type)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "fail")
			return
		}

		if webhook.Type == "charge.succeeded" {
			// TODO your code for charge
			w.WriteHeader(http.StatusOK)
		} else if webhook.Type == "refund.succeeded" {
			// TODO your code for refund
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
