package charge

var Cb_wx_pub_scan = map[string]interface{}{
	// 必填，客户端软件中展示的条码值，扫码设备扫描获取。
	"scan_code": "286801346868493272",
	// 必填，商品列表，字段解释：goods_name:商品名称，goods_num:数量。
	"goods_list": []map[string]interface{}{
		map[string]interface{}{
			"goods_name": "iPhone",
			"goods_num":  "1",
		},
		map[string]interface{}{
			"goods_name": "iPad",
			"goods_num":  "2",
		},
	},
	// 可选，指定支付方式，指定不能使用信用卡支付可设置为 no_credit 。
	"limit_pay": "no_credit",
	// 可选，终端号，要求不同终端此号码不一样，会显示在对账单中，如A01、SH008等。
	"terminal_id": "SH008",
}
