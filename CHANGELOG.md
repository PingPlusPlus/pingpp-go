## 3.0.4
*  添加身份认证功能
*  更新签名规则
*  增加批量退款、批量付款接口
*  增加报关接口
*  Refund对象新增funding_source字段
*  添加transfer更新接口

## 3.0.3
* refund对象增加Charge_order_no、Transaction_no字段
* ApiBackend.NewRequest函数使用参数key,而不是全局Key

## 3.0.2
*  统一无卡的短信(sms_code)参数(card和customer)

## 3.0.1
*  添加移动快捷支付功能

## 3.0.0
*  添加接口签名(RSA)
*  Example下的实例更新

## 2.1.5  
* refund 增加 status 字段    

## 2.1.4
* 解决 get list 时参数为空会 panic 的问题。

## 2.1.3  
* 更改 extra 字段为 map，
* 解决 bool 类型为 false 时 json 不解析的问题
* 解决长整型 json 解析时变成科学记数法的问题

##2.1.2
* 增加 transfer 对象

## 2.1.1
* 修改 summary 中两个字段

## 2.1.0
* sdk 整体重构。增加 jssdk 获取签名方法。

## 2.0.1
* 新增微信红包

## 2.0.0
* 更改：
新增渠道 bfb,wx_pub

## 1.0.1
* 更改：
Credential 字段不再一次解析完，而是作为一个interface{}对象，如果需要进一步解析，可以再次调用 Go 语言的 JSON 解析方法

### 1.0.0
* 初始发布版本
