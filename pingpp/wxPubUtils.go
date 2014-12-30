package pingpp

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Token struct {
	access_token  string
	expire_in     string
	refresh_token string
	open_id       string
	scope         string
}

func CreateOauthUrlForCode(app_id string, redirect_url string, more_info bool) (request_url string) {
	body := url.Values{}
	body.Add("app_id", app_id)
	body.Add("redirect_url", redirect_url)
	body.Add("response_type", "code")
	if more_info == true {
		body.Add("scope", "snsapi_userinfo")
	} else {
		body.Add("scope", "snsapi_base")
	}
	body.Add("state", "STATE#wechat_redirect")
	request_url = "https://open.weixin.qq.com/connect/oauth2/authorize?" + body.Encode()
	return
}

func CreateOauthUrlForOpenid(app_id string, app_secret string, code string) (request_url string) {
	body := url.Values{
		"app_id":     {app_id},
		"app_secret": {app_secret},
		"code":       {code},
		"grant_type": {"authorization_code"},
	}
	request_url = "https://api.weixin.qq.com/sns/oauth2/access_token?" + body.Encode()
	return
}

func WxPubGetWithAccessToken(app_id string, app_secret string, code string) ([]byte, error) {
	var token Token
	request_url := CreateOauthUrlForOpenid(app_id, app_secret, code)
	client := &http.Client{}
	response, err := client.Get(request_url)
	if err == nil {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resultbytes := buf.Bytes()
		err := json.Unmarshal(resultbytes, &token)
		return resultbytes, err
	} else {
		log.Printf("Http Request for OpenId Failed")
	}
	return nil, err
}
