package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pingpp/pingpp"
)

func main() {
	var myserver *http.ServeMux
	myserver = http.NewServeMux()
	myserver.HandleFunc("/webhook", webhook)
	//设置监听的端口,5623 端口只是示例使用，具体端口号根据你的代码自行定义
	err := http.ListenAndServe(":5624", myserver)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		webhook, err := pingpp.ParseWebhooks(buf.Bytes())
		if err != nil {
			fmt.Fprintf(w, "fail")
		}
		aa, _ := json.Marshal(webhook)
		fmt.Println(string(aa))
	}
}
