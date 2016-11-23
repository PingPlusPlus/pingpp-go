# Pingpp Go SDK
========


## 简介
pingpp 文件夹里是 SDK 文件

## 版本要求
建议 Go 语言版本 1.4 以上 

## 安装

导入 pingpp 模块

```go
go get github.com/pingplusplus/pingpp-go/pingpp
```

导入后，在调用的时候需要

```go
import (pingpp "github.com/pingplusplus/pingpp-go/pingpp")
```
具体使用相应模块的话还需要

```go
import (pingpp "github.com/pingplusplus/pingpp-go/pingpp/xxx")
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

### Charge查询
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

### 红包查询
```go
//查询单个 RedEnvelope 对象
redenvelope, err := redEnvelope.Get(red_id)
```

```go
//查询 RedEnvelope 列表
redenvelope, err := redEnvelope.List(&redEnvelopeListParams)
```

### event查询
```go
//查询单个 event 对象
event, err := event.Get(red_id)
```

### 身份认证
```go
//鉴别用户身份证、银行卡信息的真伪
result, err := identification.New(&identificationParams)
```

### 批量退款
```go
//发起批量退款
batch_refund, err := batchRefund.New(params)
```

### 批量企业付款
```go
//发起批量企业付款
 batch_transfer, err := batchTransfer.New(params)
```

## Debug
SDK 提供了 debug 模式。只需要更改 pingpp.go 文件中的 LogLevel 变量值，即可触发相应级别的 log，代码中对级别有注释。默认的级别是 2

## 版本号
调用

```go
pingpp.Version()
```
会返回 sdk 版本号

## 中文报错信息
Ping++ 支持中文和英文两种语言的报错信息。SDK 默认的 Accept-Language 是英文的，如果您想要接收到的错误提示是中文的，只需要设置一下即可：

```go
pingpp.AcceptLanguage = "zh-CN"
```

**详细信息请参考 [API 文档](https://pingxx.com/document/api?go)**。


