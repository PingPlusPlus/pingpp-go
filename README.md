# Pingpp Go SDK
========


## 简介
pingpp 文件夹里是 SDK 文件

## 版本要求
建议 Go 语言版本 1.3 以上 

## 安装
```
// 导入 pingpp 模块
import (pingpp "github.com/pingplusplus/pingpp-go/pingpp")
go get github.com/pingplusplus/pingpp-go/pingpp
```

## 接入方法

### 初始化
    
```go    
// 设置 API-KEY 
pingpp.Key= "YOUR-KEY"
```

### 支付
```go
//获得的第一个参数即是 Charge 对象
charge, err := charge.New(&chargeParams)
```

### 查询
```go
//查询单个 Charge 对象
charge, err := charge.Get(ch_id)
```

```go
//查询 Charge 列表
charges, err := charge.List(&chargeListParams)

```

### 退款
``` go
//charge_id为待退款的Charge的ID
refund, err := refund.New(charge_id, refundParams)
```

### 退款查询
```go
//查询单个Refund对象
refund, err := refund.Get(ch_id, re_id)
```

```go
//查询Refund对象列表
refunds, err := refund.List(ch_id, &refundListParams)
```


### 微信红包
```go
//获得的第一个参数即是 RedEnvelope 对象
redenvelope, err := redEnvelope.New(&redEnvelopeParams)
```

### 查询
```go
//查询单个 RedEnvelope 对象
redenvelope, err := redEnvelope.Get(red_id)
```

```go
//查询 RedEnvelope 列表
redenvelope, err := redEnvelope.List(&redEnvelopeListParams)
```

## Debug
SDK 提供了 debug 模式。只需要更改 pingpp.go 文件中的 LogLevel 变量值，即可触发相应级别的 log，代码中对级别有注释。默认的级别是 2

**详细信息请参考 [API 文档](https://pingxx.com/document/api?go)。**
