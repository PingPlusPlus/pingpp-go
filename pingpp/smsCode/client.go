package smsCode

import (
	"fmt"
	"log"
	"net/url"
	pingpp "pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

//查询指定 card 对象
func Get(sms_code string) (*pingpp.SmsCode, error) {
	return getC().Get(sms_code)
}

func (c Client) Get(sms_code string) (*pingpp.SmsCode, error) {
	var body *url.Values
	body = &url.Values{}
	smsCode := &pingpp.SmsCode{}
	err := c.B.Call("GET", fmt.Sprintf("/sms_codes/%v", sms_code), c.Key, body, nil, smsCode)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get smsCode error: %v\n", err)
		}
	}
	return smsCode, err
}
