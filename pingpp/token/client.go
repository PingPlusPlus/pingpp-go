package token

import (
	"fmt"
	"log"
	"net/url"
	pingpp "pingpp-go/pingpp"
	"time"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// 发送 Token 请求
func New(params *pingpp.TokenParams) (*pingpp.Token, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.TokenParams) (*pingpp.Token, error) {
	start := time.Now()
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("TokenParams Marshall Errors is : %q\n", errs)
		}
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of card request to pingpp is :\n %v\n ", string(paramsString))
	}

	token := &pingpp.Token{}
	errch := c.B.Call("POST", "/tokens", c.Key, nil, paramsString, token)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("Token completed in ", time.Since(start))
	}
	return token, errch

}

//查询指定 token 对象
func Get(tok_id string) (*pingpp.Token, error) {
	return getC().Get(tok_id)
}

func (c Client) Get(tok_id string) (*pingpp.Token, error) {
	var body *url.Values
	body = &url.Values{}
	token := &pingpp.Token{}
	err := c.B.Call("GET", fmt.Sprintf("/tokens/%v", tok_id), c.Key, body, nil, token)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Card error: %v\n", err)
		}
	}
	return token, err
}
