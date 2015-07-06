#  SDK Server 端接入文档(golang)
## 简介

### URL
Server 端需要给 Client 提供一个请求 URL 。客户端向该 URL 发起请求，服务端接收相应数据并发起支付。

### SDK 简介
Go 语言的 SDK 基于 Ping++ 的 API 进行封装，以提供对于 Go 开发者而言更简单友好的接入方法。

目前 SDK 包含 charge、refund、redEnvelope、event、transfer 和 utils 模块。
* charge 模块封装了 charge 接口，可通过该接口发起交易。
* refund 模块封装了 refund 接口，可通过该接口发起退款。
* redEnvelope 封装了微信红包的接口，可通过该接口发起微信红包。
* event 封装了 webhooks 的事件接口，可通过该接口查询 Event。
* transfer 则是微信企业付款，通过该接口可以在微信公众号中实现企业向用户付款。
* utils 是辅助的工具模块，目前只放置了微信公众帐号付款中的获取 openid 的方法和调用微信 jssdk 时获取所需签名的方法。



## 接入 SDK 
接入 Server SDK 需要以下五步，为了方便调试，GO
 SDK 提供了debug 开关。分为四个等级：
```go
 // loglevel 是 debug 模式开关.
// 0: no logging
// 1: errors only
// 2: errors + informational (default)
// 3: errors + informational + debug
var LogLevel = 2
```
由上可以看到默认是2，您可以根据自己的需求更改这个等级。例如：
```go
pingpp.LogLevel = 4
```
 说完这个，言归正传，我们看一下接入支付的具体步骤。

### Step1： 下载导入
Go SDK 支持 1.4.2 以上的 Go 语言版本，并且依赖了第三方包 simplejson。
如果项目本身没有依赖simplejson，导入时首先要获取该包：
获取 SDK ：
```go
go get -u github.com/pingplusplus/pingpp-go/pingpp
```

在项目中导入 SDK：
```go
import "github.com/pingplusplus/pingpp-go/pingpp"
```

在调用相应接口时还需要引入相应的模块，例如调用 charge 接口时，还需要导入：
```go
import "github.com/pingplusplus/pingpp-go/pingpp/charge"
```
### Step2：设置 API KEY
在 Ping++ 的管理平台可以获得 Livekey 和 Testkey。
* 如果你没有填入参数或者是想在测试模式下发起交易。设置 Testkey。
* 如果已经填写了渠道参数，想发起真实交易，设置 Livekey。
```go
pingpp.Key = "YOUR Test/Livekey"
```

### Step3：调用接口获取凭据
Server 端无论是发起付款，还是退款或者发送微信红包等，其他步骤基本一样，区别就是在这里调用相应的不同对象。这里只介绍支付所需的付款和退款。红包等其他接口的将会在相应的接入文档中单独介绍。

#### 发起付款
##### 1. 获取支付凭据
charge 模块封装了 charge 接口，供发起交易请求使用。
charge.New 方法用于发起交易请求获取支付凭据。有两个返回值，第一个是支付凭据，也就是 Charge 对象，Charge 对象的字段介绍请看 [API 文档]()。第二个返回值是错误信息 error。
以 alipay 为例展示如何发起支付请求，其他渠道如何发起支付请求详见[关于渠道]()：
```go
// metadata 是 map 类型
metadata := make(map[string]interface{})
metadata["color"] = "red"

// extra 是 map 类型,相关取值详见 api 文档
extra := make(map[string]interface{})


/**
 *@param Order_no         商户订单号，商户对订单的唯一标识。必填字段
 *@param App              appid，Ping++ 分配的应用 ID。必填字段
 *@Param Amount           订单金额，单位为分,如要表示 1 元，这里是 100。必填字段
 *@param Channel          渠道名称，如 alipay,即支付宝手机支付。必填字段
 *@param Currency         币种，目前只支持 cny，即人民币。必填字段
 *@param Client_ip        发起支付请求的 Client 的 IPv4 地址。必填字段
 *@param Extra            某些渠道所需的额外参数。可选字段
 *@param Subject          商品的标题。必填字段
 *@param Body             商品的描述信息。必填字段
 *@param Time_expire      订单失效事件， 10 位长度的时间戳。可选字段
 *@param Description      订单附加说明。可选字段
 *@param Metadata         用户可自定义的字段，格式是键值对，支持最多 10 对键值对，不支持嵌套，总长度要在 1000 个 unicode 字符。可选字段
 */
params := &pingpp.ChargeParams{
// 由于 Go 语言 json 解析的特性，所有的参数的首字母都需要大写
  Order_no:   "123456789", 
  App:         pingpp.App{Id: "app_PmfffHParTiD84aj"}, 
  Amount:      100, 
  Channel:     "alipay", 
  Currency:    "cny", 
  Client_ip:   "127.0.0.1", 
  Extra:       extra,
  Subject:     "Your Subject", 
  Body:        "Your Body", 
  Time_expire: 1410834527 
  Metadata:    metadata,
  Description: "Your Description",
}

//返回的第一个参数 ch 是获取到的 Charge 对象，你需要将其转换成 json 给 Client ，或者 Client 接收后转换。第二个返回值 err 是错误信息，若正确返回 Charge 对象，则为 nil
ch, err := charge.New(params)
if err != nil {
  log.Fatal(err)
}
```

如果正常得到支付凭据，即 Charge 对象，示例如下：
```json
{
  "id": "ch_Hm5uTSifDOuTy9iLeLPSurrD",
  "object": "charge",
  "created": 1425095528,
  "livemode": true,
  "paid": false,
  "refunded": false,
  "app": "app_1234567890abcDEF",
  "channel": "alipay",
  "order_no": "123456789",
  "client_ip": "127.0.0.1",
  "amount": 100,
  "amount_settle": 0,
  "currency": "cny",
  "subject": "Your Subject",
  "body": "Your Body",
  "extra":{},
  "time_paid": null,
  "time_expire": 1425181928,
  "time_settle": null,
  "transaction_no": null,
  "refunds": {
    "object": "list",
    "url": "/v1/charges/ch_Hm5uTSifDOuTy9iLeLPSurrD/refunds",
    "has_more": false,
    "data": []
  },
  "amount_refunded": 0,
  "failure_code": null,
  "failure_msg": null,
  metadata": {
	  "color":"red"
	 },
  "credential": {
    "object": "credential",
    "alipay":{
      "orderInfo": "_input_charset=\"utf-8\"&body=\"tsttest\"&it_b_pay=\"1440m\"¬ify_url=\"https%3A%2F%2Fapi.pingxx.com%2Fnotify%2Fcharges%2Fch_jH8uD0aLyzHG9Oiz5OKOeHu9\"&out_trade_no=\"1234dsf7uyttbj\"&partner=\"2008451959385940\"&payment_type=\"1\"&seller_id=\"2008451959385940\"&service=\"mobile.securitypay.pay\"&subject=\"test\"&total_fee=\"1.23\"&sign=\"dkxTeVhMMHV2dlRPNWl6WHI5cm56THVI\"&sign_type=\"RSA\""
    }
  },
  "description": null
}
```
##### 2. 将支付凭据返回给 Client 
付款需要将获得的支付凭据发送给 Client ， Client 接收支付凭据调起支付控件完成交易。具体传递需要商户根据自身处理方式去传递。由于 Go  SDK 获得的是 Charge 对象，需要将其转换成 JSON 再传给 Client 。
```go
chString, _ := json.Marshal(ch)
fmt.Fprintln(w, string(chString))
```

#### 发起退款
发起退款只在 Server 端 调用 SDK 发起即可，不需要客户端 SDK 参与。
Go SDK 的 refund 模块封装了 refund 接口，供发起退款请求使用。
refund.New 方法用于发起交易请求获取支付凭据。有两个返回值，第一个是 refund 对象，refund 对象的字段介绍请看 [API 文档]()。第二个返回值是错误信息 error。
以 alipay 为例展示如何发起退款请求：
```go
metadata := make(map[string]interface{})
metadata["color"] = "red"
/**
 *@Param Amount           退款金额，需要小于等于付款金额，单位为分。可选字段，不填的话默认全额退款 
 *@param Description      退款描述，必填字段
 *@param Metadata         用户可自定义的字段，格式是键值对，支持最多 10 对键值对，不支持嵌套，总长度要在 1000 个 unicode 字符。可选字段
 */
params := &pingpp.RefundParams{
	Amount:      100, 
	Description: "Refund Description", 
	Metadata:    metadata,
}
//当 err 为空时，re 即为获得的 Refund 对象
re, err := refund.New("ch_Xsr7u35O3m1Gw4ed2ODmi4Lw", params) //第一个参数为需要退款的订单的 id 
if err != nil {
	log.Fatal(err)
}
```

完成退款请求后会获得退款对象：
```json
{
   "id": "re_SG0mnjTD3jAHimbvDKjnXLC9",
   "object": "refund",
   "order_no": "SG0mnjTD3jAHimbvDKjnXLC9",
   "amount": 100,
   "created": 1427555346,
   "succeed": false,
   "time_succeed": 1427555348,
   "description": "Refund Description",
   "failure_code": null,
   "failure_msg": null,
   metadata": {
	  "color":"red"
	 },
   "charge": "ch_Xsr7u35O3m1Gw4ed2ODmi4Lw"
}
```

获得退款对象表明发起退款请求成功，渠道接收了退款请求并且进行相应处理。
不同的渠道处理退款的时间不一致，所以收到 Webhooks 通知的时间与渠道相关。关于退款不同渠道的区别，详见[关于渠道]()。
退款是否成功是根据 Refund 对象中的 succeed 字段是否为 true 确定。商户需要通过接收 Webhooks 通知来判断退款状态。

### Step5：接收 webhooks 通知
用户在 Client 完成支付时，渠道会直接返回给 Client 支付结果。与此同时，Ping++ 会给商户指定的 Server 发送 Webhooks 通知。
通知以 POST 方式主动发送给商户 Server ，内容是 Event 事件。Event 事件的字段描述请参考 [API 文档]()。
商户需要以 Webhooks 通知作为支付成功与否的判断依据，不能使用 Client 的通知来判断支付结果。
#### 1. 配置 Webhooks
在 Ping++ 的平台上配置 Webhooks。商户将指定的 Server 地址配置到 Webhooks 里，同时还需要指定接收的事件类型，事件所属的支付环境（Test 环境/ Live 环境）等。如何配置 Webhooks 请参考 [Webhooks 使用指南]()。
Event 对象中的 type 字段表示事件类型。支付结果的事件类型是 charge.succeeded，该事件只在付款成功后才被发送；退款的事件类型是 refund.succeeded，同理只在退款成功后才会发送该事件。
商户需要选择接收这两个事件类型用来判断支付或退款是否成功。
#### 2. 获取 Webhooks 通知
支付成功时，Ping++ 主动以 POST 的方式发送 Webhooks 通知给商户指定的 Server 地址。商户需要在该地址监听并接收 Webhooks 通知。
* 接收到 Webhooks 通知后需要返回给 Ping++  Server 状态 2xx(建议 200)。否则，建议返回状态码 500 表示未收到。
* 若商户 Server 未正确返回状态码 2xx，Ping++ 会在 25 小时内向商户 Server 最多发送 8 次 Webhooks 通知。时间间隔分别为 2min、10min、10min、1h、2h、6h、15h，直到用户向 Ping++ 返回 Server 状态 2xx 或者超过最大重发次数为止。

如下举例说明 Go 如何接收 Webhooks ，具体如何接收、处理需要商户根据自己的工程确定。不过 SDK 里提供了解析 Webhooks 的方法，放在 webhook.go 文件里，方法为 ParseWebhooks(webhooks []byte) (*Event, error) 。有两个返回值，第一个是解析后的 Event 对象，第二个是 error。 

```go
func webhook(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
    signature := r.Header.Get("x-pingplusplus-signature") //从 header 获取签名字段
    webhook, err := pingpp.ParseWebhooks(buf.Bytes()) // 调用parseWebhooks 方法解析 event 对象 
    if err != nil {
      fmt.Fprintf(w, "fail")
    } else {
      if webhook.Type == "charge.succeeded" {//由于 Go 解析 json 的特性，Event 对象中所有字段在取的时候首字母都需大写
        w.WriteHeader(http.StatusOK) // 获取到指定的事件后返回状态码200
      } else if webhook.Type == "refund.succeeded" {
        w.WriteHeader(http.StatusOK) // 获取到指定的事件后返回状态码200
      } else {
        w.WriteHeader(http.StatusInternalServerError) // 没有获取到通知返回 Server 状态500      
      }
    }
  }
}
```

**附**：类型为 charge.succeeded 的事件示例如下：
```json
{
    "id": "evt_ugB6x3K43D16wXCcqbplWAJo",
    "created": 1427555101,
    "livemode": true,
    "type": "charge.succeeded",
    "data": {
        "object": {
            "id": "ch_Xsr7u35O3m1Gw4ed2ODmi4Lw",
            "object": "charge",
            "created": 1427555076,
            "livemode": true,
            "paid": true,
            "refunded": false,
            "app": "app_1234567890abcDEF",
            "channel": "upacp",
            "order_no": "123456789",
            "client_ip": "127.0.0.1",
            "amount": 100,
            "amount_settle": 0,
            "currency": "cny",
            "subject": "Your Subject",
            "body": "Your Body",
            "extra": {},
            "time_paid": 1427555101,
            "time_expire": 1427641476,
            "time_settle": null,
            "transaction_no": "1224524301201505066067849274",
            "refunds": {
                "object": "list",
                "url": "/v1/charges/ch_L8qn10mLmr1GS8e5OODmHaL4/refunds",
                "has_more": false,
                "data": []
            },
            "amount_refunded": 0,
            "failure_code": null,
            "failure_msg": null,
            "metadata": {},
            "credential": {},
            "description": null
        }
    },
    "object": "event",
    "pending_webhooks": 0,
    "request": "iar_qH4y1KbTy5eLGm1uHSTS00s"
}
```
需要注意的是 event 的 data 字段中包含了事件类型所属的对象，例如 charge.succeeded 中 data 里是 Charge 对象。其 key 是 object 字段。商户接收到即该事件即表明付款成功，也可从 charge 对象中判断 paid 字段是否为 true 来确定交易是否成功。

类型为 refund.succeeded 的事件示例如下：
```json
{
    "id": "evt_gJKelawq06CiPojS5gt3noQA",
    "created": 1427555348,
    "livemode": true,
    "type": "refund.succeeded",
    "data": {
        "object": {
            "id": "re_SG0mnjTD3jAHimbvDKjnXLC9",
            "object": "refund",
            "order_no": "SG0mnjTD3jAHimbvDKjnXLC9",
            "amount": 100,
            "created": 1427555346,
            "succeed": true,
            "time_succeed": 1427555348,
            "description": "Refund Description",
            "failure_code": null,
            "failure_msg": null,
            "metadata": {},
            "charge": "ch_Xsr7u35O3m1Gw4ed2ODmi4Lw"
        }
    },
    "object": "event",
    "pending_webhooks": 0,
    "request": "iar_Ca1Oe10OqTSOPOmzX9Hi1a5"
}
```
event 的 data 字段中包含了事件类型所属的对象，例如 refund.succeeded 中 data 里是 Refund 对象。其 key 是 object 字段。商户接收到即该事件即表明付款成功，也可从 refund 对象中判断 succeed 字段是否为 true 来确定交易是否成功。


#### 验证 Webhooks 通知
Ping++ 对 Webhooks 通知进行了 RSA 签名，商户可使用签名验证 Webhooks 通知的合法性。
* 签名放在 Header 的自定义字段 x-pingplusplus-signature 中。
* 签名用 RSA 私钥对 Webhooks 通知使用 RSA-SHA256 算法进行签名，以 base64 格式输出。

需要注意的是，在 Test 环境（用 Testkey 发起交易）不会生成签名。可以从 Header 中取到 x-pingplusplus-signature 字段，但是值为空。只有在 Live 环境（用Livekey 发起交易）才会生成签名。

##### 验证签名
Ping++ 在管理平台中提供了 RSA 公钥，供验证签名。公司签约完成后，公钥信息在管理平台的「公司信息」中获取。验证签名需要以下步：
1. 从 header 取出签名并对其进行 base64 解码
```go
signature := r.Header.Get("x-pingplusplus-signature")
```
2. 把 Webhooks 通知、Ping++ 管理平台提供的 RSA 公钥、 和 base64 解码后的签名三者一同放入 RSA 的签名函数中进行非对称的签名运算,来判断签名是否验证通过。

由于某些原因，验证签名的方法并没有集成在 SDK 里，但是我们提供了 Demo ，具体参考[Go SDK Demo](https://github.com/PingPlusPlus/pingpp-go/blob/master/verifyDemo.go)。


**当商户 Server 能够正确获得支付凭据并且正确返回给 Client ，而且也完成了对 Webhooks 的监听、接收和验证。至此， Server  SDK 的接入已经完成。现在需要去完成 Client SDK 的接入**