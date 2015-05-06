package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"pingpp/pingpp"
)

func main() {
	var myserver *http.ServeMux
	myserver = http.NewServeMux()
	myserver.HandleFunc("/notify", notify)
	//设置监听的端口,5623 端口只是示例使用，具体端口号根据你的代码自行定义
	err := http.ListenAndServe(":5624", myserver)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func notify(w http.ResponseWriter, r *http.Request) {
	// var charge pingpp.Charge
	// var refund pingpp.Refund
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		ch, re, err := pingpp.ParseNotify(buf.Bytes())
		fmt.Println(ch)
		if err != nil {
			//your code
			fmt.Fprintf(w, "fail")
		} else if ch != nil {
			if ch.Paid == true {
				//your code
				fmt.Fprintf(w, "success")
			} else {
				fmt.Fprintf(w, "fail")
			}
		} else if re != nil {
			if re.Succeed == true {
				fmt.Fprintf(w, "success")
			} else {
				fmt.Fprintf(w, "fail")
			}
		} else {
			fmt.Fprintf(w, "fail")
		}
	}
}
