package charge

var Cb_alipay = map[string]interface{}{
	// 可选，支付类型。默认值为：1（商品购买）。
	"payment_type": "1",
	// 可选，分账列表。
	"split_fund_info": []map[string]interface{}{
		map[string]interface{}{
			"account": "2088866088886666", // 接受分账资金的支付宝账户ID。
			"amount":  1,                  // 分账的金额。
			"desc":    "split_desc desc",  // 分账描述信息。
		},
	},
}
