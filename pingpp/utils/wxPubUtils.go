package utils

import (
	"bytes"
	"crypto/sha1"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"log"
	"net/http"
	"net/url"
)

type Token struct {
	Access_token  string `json:"access_token"`
	Expires_in    string `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Openid        string `json:"openid"`
	Scope         string `json:"scope"`
}

type AccessToken struct {
	Access_token string `json:"access_token"`
	Expires_in   uint64 `json:"expires_in"`
}

type JsapiTicket struct {
	Errcode    uint   `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Ticket     string `json:"ticket"`
	Expires_in uint   `json:"expires_in"`
}

type Wx_pub struct {
	AppId     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	TimeStamp string `json:"timeStamp"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

/**
 * 用于获取授权 code 的 URL 地址，此地址用于用户身份鉴权，获取用户身份信息，同时重定向到 $redirect_uri
 * @param $app_id 微信公众号应用唯一标识
 * @param $redirect_uri 授权后重定向的回调链接地址，重定向后此地址将带有授权code参数，
 *                      该地址的域名需在微信公众号平台上进行设置，
 *                      步骤为：登陆微信公众号平台 => 开发者中心 => 网页授权获取用户基本信息 => 修改
 * @param bool $more_info FALSE 不弹出授权页面,直接跳转,这个只能拿到用户 openid
 *                        TRUE 弹出授权页面,这个可以通过 openid 拿到昵称、性别、所在地，
 * @return string 用于获取授权 code 的 URL 地址
 */
func CreateOauthUrlForCode(app_id string, redirect_uri string, more_info bool) (request_url string) {
	body := url.Values{}
	body.Add("appid", app_id)
	body.Add("redirect_uri", redirect_uri)
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

/**
 * 获取openid的URL地址
 * @param $app_id 微信公众号应用唯一标识
 * @param $app_secret 微信公众号应用密钥（注意保密）
 * @param $code 授权 code, 通过调用 WxpubOAuth::createOauthUrlForCode 来获取
 * @return string 获取 openid 的 URL 地址
 */
func CreateOauthUrlForOpenid(app_id string, app_secret string, code string) (request_url string) {
	body := url.Values{
		"appid":      {app_id},
		"secret":     {app_secret},
		"code":       {code},
		"grant_type": {"authorization_code"},
	}
	request_url = "https://api.weixin.qq.com/sns/oauth2/access_token?" + body.Encode()
	return
}

/**获取微信公众号授权用户唯一标识
:param app_id: 微信公众号应用唯一标识
:param app_secret: 微信公众号应用密钥（注意保密）
:param code: 授权 code, 通过调用 WxpubOAuth.createOauthUrlForCode 来获取
:return: openid 微信公众号授权用户唯一标识, 可用于微信网页内支付
*/

func GetOpenId(app_id string, app_secret string, code string) ([]byte, error) {
	var token Token
	request_url := CreateOauthUrlForOpenid(app_id, app_secret, code)
	client := &http.Client{}
	response, err := client.Get(request_url)
	if err == nil {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resultbytes := buf.Bytes()
		err := pingpp.JsonDecode(resultbytes, &token)
		return resultbytes, err
	} else {
		log.Printf("Http Request for OpenId Failed")
	}
	return nil, err
}

/**
 * 获取微信公众号 jsapi_ticket
 * @param $app_id 微信公众号应用唯一标识
 * @param $app_secret 微信公众号应用密钥（注意保密）
 * @return array 包含 jsapi_ticket 的数组或者错误信息
 */
func GetJsapiTicket(app_id string, app_secret string) (ticket string) {
	var accessToken AccessToken
	var jsapiTicket JsapiTicket
	body := url.Values{}
	body.Add("appid", app_id)
	body.Add("secret", app_secret)
	body.Add("grant_type", "client_credential")
	accessTokenUrl := "https://api.weixin.qq.com/cgi-bin/token?" + body.Encode()
	client := &http.Client{}
	response, err := client.Get(accessTokenUrl)
	if err == nil {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resultbytes := buf.Bytes()
		errs := pingpp.JsonDecode(resultbytes, &accessToken)
		if pingpp.LogLevel > 2 {
			log.Printf("response from wx for AccessToken struct is :%v\n", jsapiTicket)
		}
		if errs != nil {
			if pingpp.LogLevel > 0 {
				log.Printf("Cannot Unmarshal AccessToken struct:", errs)
			}
		}
	} else {
		if pingpp.LogLevel > 0 {
			log.Printf("Http Request for AccessToken Failed")
		}
	}

	body = url.Values{}
	body.Add("access_token", accessToken.Access_token)
	body.Add("type", "jsapi")
	jsapiTicketUrl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?" + body.Encode()
	client = &http.Client{}
	response, err = client.Get(jsapiTicketUrl)

	if err == nil {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resultbytes := buf.Bytes()
		err := pingpp.JsonDecode(resultbytes, &jsapiTicket)
		if pingpp.LogLevel > 2 {
			log.Printf("response from wx for jsapiTcicket struct is :%v\n", jsapiTicket)
		}
		if err != nil {
			if pingpp.LogLevel > 0 {
				log.Printf("Cannot Unmarshal JsapiTicket struct:", err)
			}
		}
		ticket = jsapiTicket.Ticket
		return ticket
	} else {
		log.Printf("Http Request for JsapiTicket Failed")
	}
	return
}

/**
 * 生成微信公众号 js sdk signature
 * @param charge Charge
 * @param jsapi_ticket string
 * @param urls string  当前页面的 url， 必须要动态获取
 * @return signatrue []byte
 */
func GetSignature(charge *pingpp.Charge, jsapi_ticket string, urls string) (signatrue []byte) {
	var wx Wx_pub
	a := charge.Credential
	m := a["wx_pub"]
	s, _ := pingpp.JsonEncode(m)
	pingpp.JsonDecode(s, &wx)
	jsapi_tickets := "jsapi_ticket=" + jsapi_ticket
	nonce_str := "noncestr=" + wx.NonceStr
	time_stamp := "timestamp=" + wx.TimeStamp
	urls = "url=" + urls
	signs := jsapi_tickets + "&" + nonce_str + "&" + time_stamp + "&" + urls
	h := sha1.New()
	h.Write([]byte(signs))
	signatrue = h.Sum(nil)
	return signatrue
}
