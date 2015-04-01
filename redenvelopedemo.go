package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pingpp-go/pingpp"
)

var mymux *http.ServeMux

func Run() {

	//异步通知配置
	// pingpp.Notify("/notify", ":8080", 100)
	// asyn()

	//服务器路由
	mymux = http.NewServeMux()
	bind()

	err := http.ListenAndServe(":1281", mymux) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
	bind the router
*/
func bind() {
	mymux.HandleFunc("/redenvelope", redenvelope)
}

func redenvelope(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.NotFound(w, r)
	} else if r.Method == "POST" {
		// var jsonstring string
		// jsonstring = `{"order_no": "1234567890111222","amount": 1,"app": {"id":"YourApp"},"channel": "alipay","currency": "cny","client_ip": "192.168.1.1","subject": "test-subject","body": "test-body","metadata": {"color": "red"}}`
		var redenvelopeParams pingpp.RedEnvelopeParams
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Println(buf.String())
		json.Unmarshal(buf.Bytes(), &redenvelopeParams)
		redenvelopeClient := pingpp.GetRedEnvelopeClient("YOUR-KEY")
		redenvelope, err := redenvelopeClient.New(&redenvelopeParams)
		if err != nil {
			errorbytes, _ := json.Marshal(err)
			fmt.Fprint(w, string(errorbytes))
		} else {
			redenvelopebytes, _ := json.Marshal(redenvelope)
			fmt.Fprint(w, string(redenvelopebytes))
		}

	}
	return
}

/*
	定义自己的notify行为
*/
// func asyn() {
// 	go func() {
// 		for {
// 			charge := <-pingpp.ChargeChan
// 			//TO-DO
// 			fmt.Printf("notify charge: %v \n", charge)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			refund := <-pingpp.RefundChan
// 			//TO-DO
// 			fmt.Printf("notify charge: %v \n", refund)
// 		}
// 	}()
// }

func main() {
	Run()
}
