package wxlite

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

const maxBytes int64 = 10 * 1024 * 1024

// Client wxlite 请求客户端
var httpclient = http.Client{Timeout: 12 * time.Second}

// GetOpenid 发送 https://api.weixin.qq.com/sns/jscode2session 请求
func GetOpenid(params *pingpp.OpenidParams) (*pingpp.Openid, error) {
	start := time.Now()
	values := make(url.Values)
	values.Add("appid", params.AppID)
	values.Add("secret", params.AppSecret)
	values.Add("js_code", params.Code)
	values.Add("grant_type", "authorization_code")
	resp, err := httpclient.Get("https://api.weixin.qq.com/sns/jscode2session?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &io.LimitedReader{R: resp.Body, N: maxBytes}
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	if pingpp.LogLevel > 2 {
		log.Println("GetOpenid completed in ", time.Since(start))
	}
	openid := &pingpp.Openid{}
	err = json.Unmarshal(bytes, openid)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return openid, nil
}
