package common

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pingplusplus/pingpp-go/demo/common/charge"
	"github.com/pingplusplus/pingpp-go/demo/common/transfer"
	"github.com/pingplusplus/pingpp-go/demo/common/withdrawal"
)

func Response(data interface{}, err error) {
	if err != nil {
		log.Fatalln("response error:", err)
		return
	}
	PrintResponse(data)
}

func PrintResponse(data interface{}) {
	content, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(content))
}

type extra struct {
	ChargeExtra   map[string]map[string]interface{}
	TransferExtra map[string]map[string]interface{}
	WithdrawExtra map[string]map[string]interface{}
}

var Extra = extra{
	ChargeExtra: ChargeExtra,
}

var ChargeExtra = map[string]map[string]interface{}{
	"alipay":           charge.Alipay,
	"alipay_pc_direct": charge.Alipay_pc_direct,
	"alipay_scan":      charge.Alipay_scan,
	"alipay_wap":       charge.Alipay_wap,
	"applepay_upacp":   charge.Applepay_upacp,
	"balance":          charge.Balance,
	"bfb_wap":          charge.Bfb_wap,
	"cb_alipay":        charge.Cb_alipay,
	"cb_wx":            charge.Cb_wx,
	"cb_wx_pub":        charge.Cb_wx_pub,
	"cb_wx_pub_qr":     charge.Cb_wx_pub_qr,
	"cb_wx_pub_scan":   charge.Cb_wx_pub_scan,
	"cmb_wallet":       charge.Cmb_wallet,
	"fqlpay_wap":       charge.Fqlpay_wap,
	"isv_qr":           charge.Isv_qr,
	"isv_scan":         charge.Isv_scan,
	"isv_wap":          charge.Isv_wap,
	"jdpay_wap":        charge.Jdpay_wap,
	"mmdpay_wap":       charge.Mmdpay_wap,
	"qgbc_wap":         charge.Qgbc_wap,
	"qpay":             charge.Qpay,
	"upacp":            charge.Upacp,
	"upacp_pc":         charge.Upacp_pc,
	"upacp_wap":        charge.Upacp_wap,
	"wx":               charge.Wx,
	"wx_lite":          charge.Wx_lite,
	"wx_pub":           charge.Wx_pub,
	"wx_pub_qr":        charge.Wx_pub_qr,
	"wx_pub_scan":      charge.Wx_pub_scan,
	"wx_wap":           charge.Wx_wap,
	"yeepay_wap":       charge.Yeepay_wap,
}

var TransferExtra = map[string]map[string]interface{}{
	"alipay":   transfer.Alipay,
	"allinpay": transfer.Allinpay,
	"balance":  transfer.Balance,
	"jdpay":    transfer.Jdpay,
	"unionpay": transfer.Unionpay,
	"wx_pub":   transfer.Wx_pub,
}

var WithdrawExtra = map[string]map[string]interface{}{
	"alipay":   withdrawal.Alipay,
	"allinpay": withdrawal.Allinpay,
	"jdpay":    withdrawal.Jdpay,
	"unionpay": withdrawal.Unionpay,
	"wx_pub":   withdrawal.Wx_pub,
}
