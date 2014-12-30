Pingpp Go SDK
=================

****

## 简介

pingpp 文件夹下是 Go SDK 文件。
以 test 结尾的 Go 文件是单元测试文件，同时也可以作为参考模板

## 版本要求

Go 语言版本建议 1.3 以上

## 接入方法

关于如何使用 SDK 请参考 [技术文档](https://pingplusplus.com/document) 或参考 pingpp 目录下后缀为 test 的 Go 文件

## 更新日志
### 2.0.0
* 更改：
新增渠道 bfb,wxpub


### 1.0.1
* 更改：
Credential 字段不再一次解析完，而是作为一个interface{}对象，如果需要进一步解析，可以再次调用 Go 语言的 JSON 解析方法

### 1.0.0
* 初始发布版本
