package pingpp

import (
	"bytes"
	"encoding/json"

	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	//"encoding/base64"
	"encoding/pem"
	"errors"
)

// 简单封装了json的Marshal功能
// pingpp.JsonEncode(param1)
func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(&v)
}

// 简单封装了json的UnMarshal功能
// Example pingpp.JsonDecode(param1, param2)
// param1：需要转换成结构体的json数据
// param2：转换后数据容器
func JsonDecode(p []byte, v interface{}) error {
	obj := json.NewDecoder(bytes.NewBuffer(p))
	obj.UseNumber()
	return obj.Decode(&v)
}

// 转换webhooks api请求的body到已定的数据结构
func ParseWebhooks(webhooks []byte) (*Event, error) {
	var event Event
	if webhooks != nil && len(webhooks) > 0 {
		err := JsonDecode(webhooks, &event)
		if err != nil {
			return nil, err
		}
	}
	return &event, nil
}

//用商户的私钥去生成签名目前在创建订单的时候使用
func GenSign(data []byte, privateKey []byte) (sign []byte, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, priv, hashFunc, hashed)
}

//用ping++公钥去验证签名目前在Webhook时候使用
func Verify(data []byte, publicKey []byte, sign []byte) (err error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, hashFunc, hashed, sign)
}
