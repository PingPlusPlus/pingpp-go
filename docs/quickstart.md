# quickstart

本文档旨在帮助你快速运行 demo 代码，以熟悉 go 的调用方式。

## 示例代码文件
pay.go 和 paytest.go 是示例代码。
**注意**
sdk 使用了 simple-json 包，需要额外导入该包。

* pay.go 里面给出了所有方法的调用示例。
* paytest.go 新建了路由去接收客户端传递的数据，然后调用了charge.New 方法。

## pay.go

1. 在 init 方法里面填上你自己的 livekey (Ping++ 分配给你的以 	
sk_live_ 开头的字符串) 或者 testkey(Ping++ 分配给你的以 	
sk_test_ 开头的字符串)。
2. 如果你调用的是响应的 New 方法，需要在相应的 params 里设置 App，如下：
```go
 App:       pingpp.App{Id: "YOUR-APP-ID"}
```
 替换 YOUR-APP-ID 为你的应用 ID (ping++ 分配给你的以 app_ 开头的字符串)。
 3. 如果调用的是 响应的 Get 方法，需要在里面填入享用的订单号，ch_id 就是你的交易的订单号，re_id 是指定订单的退款订单号，red_id 指的是红包订单号。
4. 在 main 方法里调用相应的你需要调用的方法。
5. 现在进入相应目录 go run pay.go 即可。


## paytest.go

1. 找到设置key的地方设置Key:
```go
pingpp.Key = "YOUR-KEY"
```
2. demo 里的路由是：/pay，端口1281，你可根据自己的需求设置路由
3. demo 里面解析了两个字段 amount 和 channel。而这些字段是 chargeParams 里的，根据你的客户端传递的来解析，在调用 charge.New 方法时填入即可。
4. 进入相应目录 go run paytest.go
5. 客户端根据指定的路由传递数据过去（demo解析的是 json 格式，这里的格式由你根据自己传递的格式以相应方式读取即可）
4. 客户端按照定义的路由
