## 3.2.4

**新增**
1. 分账
2. 分账明细
3. 分账接收方
4. 银行支行查询

## 3.2.3

**修改**
1. 修复了重试存在的一个 bug

## 3.2.2

**新增**
1. 签约接口
2. wxlite GetOpenid 小程序获取 openid 接口

## 3.2.1

**新增**
1. 新增 charge 渠道demo alipay_scan、wx_pub_scan、cb_alipay、cb_wx、cb_wx_pub、cb_wx_pub_qr、cb_wx_pub_scan
2. coupon 增加字段 `user_times_circulated`, order pay接口增加请求参数 `time_expire`,order 对象去除字段 `user_from`
3. token,customer,card接口下线

**修改**

## 3.2.0

**新增**
1. 合并账户系统相关接口

**修改**
1. 自动重试机制

## 3.1.1(2016-12-26)
*  添加创建、查询、更新用户接口
*  添加优惠券&优惠券模板创建、更新、删除、查询接口
*  添加充值，订单的创建、支付、取消、查询、退款接口
*  添加asset\_transaction&balance\_transaction查询接口
*  添加余额转账、提现接口

## 3.1.0 （2016-12-26）
**新增**
*  添加身份认证功能
*  更新签名规则
*  增加批量退款、批量付款接口
*  增加报关接口
*  添加transfer更新接口

**修改**
*  Refund对象新增funding_source字段
*  查询charge列表时app[id]为必填参数

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
