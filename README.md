# Pingpp Go SDK
========


## 简介
pingpp 文件夹里是 SDK 文件，不过其中 的 `charge_test.go`、`redEnvelope_test.go`和 `refund_test.go` 作为测试代码，也可以作为调用实例的参考使用。

## 版本要求
建议 Go 语言版本 1.3 以上 

## 安装
```
// 导入 pingpp 模块
import (pingpp "github.com/pingplusplus/pingpp-go")
go get pingpp
```

## 接入方法

### 初始化
    
```go    
// 设置 API-KEY 并获取 Client
client := getChargeClient(YourKey)
client := getRefundClient(YourKey)
client := GetRedEnvelopeClient(YourKey)
```

### 支付
```go
//获得的第一个参数即是 Charge 对象
charge, err := client.new(&chargeParams)
```

### 查询
```go
//查询单个 Charge 对象
charge, err := client.get(id)
```

```go
//查询 Charge 列表
charges, err := client.list(params)

```

### 退款
``` go
//charge_id为待退款的Charge的ID
refund, err := client.new(refundParams, charge_id)
```

### 退款查询
```go
//查询单个Refund对象
refund, err := client.get(charge_id, refund_id)
```

```go
//查询Refund对象列表
refunds, err := client.list(charge_id, limit, starting_after, ending_before)
```


### 微信红包
```go
//获得的第一个参数即是 RedEnvelope 对象
redenvelope, err := client.new(&redEnvelopeParams)
```

### 查询
```go
//查询单个 RedEnvelope 对象
redenvelope, err := client.get(RED-ID)
```

```go
//查询 RedEnvelope 列表
redenvelope, err := client.list(params)
```

**详细信息请参考 [API 文档](https://pingxx.com/document/api?go)。**
