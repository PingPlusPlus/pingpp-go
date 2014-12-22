package pingpp

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
	mymux.HandleFunc("/pay", pay)
}

func pay(w http.ResponseWriter, r *http.Request) {
	var chargeParams ChargeParams
	w.Header().Set("Access-Control-Allow-Origin", "*")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err := json.Unmarshal(buf.Bytes(), &chargeParams)
	if err != nil {
		chargeClient := GetChargeClient("key")
		charge, _ := chargeClient.New(&chargeParams)
		chargebytes, _ := json.Marshal(charge)
		w.Write(chargebytes)
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
