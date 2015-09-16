package main

import (
	"bytes"
	// "encoding/json"
	"fmt"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/http"
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
		// signature := r.Header.Get("x-pingplusplus-signature")
		webhook, err := pingpp.ParseWebhooks(buf.Bytes())
		fmt.Println(webhook.Type)
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
