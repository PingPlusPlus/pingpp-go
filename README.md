pingpp-go
=========


# 安装

	// 导入pingpp模块
    import (pingpp "github.com/pingplusplus/pingpp-go")
    go get pingpp

# 使用
#### 在接口调用之前，需执行如下代码：
    
    
    // 设置API-KEY并获取Client
    client := getChargeClient(YourKey)
    client := getRefundClient(YourKey)


## 创建Charge对象
	//Charge对象信息
	chargeParams := &ChargeParams{
		order_no:  "88888887",
		appid:     "app_mHarHK4KajnDWDW9",
		channel:   "alipay",
		amount:    1000,
		currency:  "cny",
		client_ip: "127.0.0.1",
		subject:   "test",
		body:      "bodysample",
	}
	
	//获得的第一个参数即是Charge对象
	charge, err := client.new(chargeParams)
    
    
## 查询 Charge 对象
	//查询Charge列表
    charges, err := client.list(params)
    
    //查询单个Charge对象
    charge, err := client.get(id)
    
## 创建 Refund 对象
    refundParams := &RefundParams{
		Amount:      1, //退款数量，单位为分
		Description: "Some Description",
	}
	//charge_id为待退款的Charge的ID
	refund, err := client.new(refundParams, charge_id)
	
    
## 查询 Refund 对象
	//查询Refund对象列表
    refunds, err := client.list(charge_id, limit, starting_after, ending_before)
    
    //查询单个Refund对象
    refund, err := client.get(charge_id, refund_id)
    
    
