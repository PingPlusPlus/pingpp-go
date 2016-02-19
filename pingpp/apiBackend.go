package pingpp

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//api相关的后端类型，此类型线程不安全
type ApiBackend struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// 后端处理请求方法
func (s ApiBackend) Call(method, path, key string, form *url.Values, params []byte, v interface{}) error {
	var body io.Reader
	if strings.ToUpper(method) == "POST" {
		body = bytes.NewBuffer(params)
	}

	if strings.ToUpper(method) == "GET" {
		if form != nil && len(*form) > 0 {
			data := form.Encode()
			path += "?" + data
		}
	}

	req, err := s.NewRequest(method, path, key, "application/json", body, params)

	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// 建立http请求对象
func (s *ApiBackend) NewRequest(method, path, key, contentType string, body io.Reader, params []byte) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = s.URL + path
	req, err := http.NewRequest(method, path, body)
	if LogLevel > 2 {
		log.Printf("Request to pingpp is : \n %v\n", req)
	}

	if err != nil {
		if LogLevel > 0 {
			log.Printf("Cannot create pingpp request: %v\n", err)
		}
		return nil, err
	}

	if strings.ToUpper(method) == "POST" && len(AccountPrivateKey) > 0 {
		sign, err := GenSign([]byte(params), []byte(AccountPrivateKey))
		if err != nil {
			if LogLevel > 0 {
				log.Printf("Cannot create RSA signature: %v\n", err)
			}
			return nil, err
		}
		encodeSign := base64.StdEncoding.EncodeToString(sign)
		req.Header.Add("Pingplusplus-Signature", encodeSign)
	}

	// 添加Auth参数(获取参见 https://www.pingxx.com/guidance/)满足ping++ api的http BasicAuth验证
	req.SetBasicAuth(Key, "")
	req.Header.Add("Pingpluspplus-Version", apiVersion)
	req.Header.Add("User-Agent", "Pingpp go SDK version:"+Version())
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept-Language", AcceptLanguage)

	return req, nil
}

// 处理http请求
func (s *ApiBackend) Do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}
	start := time.Now()
	res, err := s.HTTPClient.Do(req)

	if LogLevel > 0 {
		log.Printf("Request to pingpp completed in %v\n", time.Since(start))
	}

	if err != nil {
		if LogLevel > 0 {
			log.Printf("Request to Pingpp failed: %v\n", err)
		}
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if LogLevel > 0 {
			log.Printf("Cannot parse Pingpp response: %v\n", err)
		}
		return err
	}

	if res.StatusCode >= 400 {
		var errMap map[string]interface{}
		JsonDecode(resBody, &errMap)

		if e, ok := errMap["error"]; !ok {
			err := errors.New(string(resBody))
			if LogLevel > 0 {
				log.Printf("Unparsable error returned from Pingpp: %v\n", err)
			}
			return err
		} else {
			root := e.(map[string]interface{})
			err := &Error{
				Type:           ErrorType(root["type"].(string)),
				Msg:            root["message"].(string),
				HTTPStatusCode: res.StatusCode,
			}

			if code, found := root["code"]; found {
				err.Code = ErrorCode(code.(string))
			}

			if param, found := root["param"]; found {
				err.Param = param.(string)
			}

			if LogLevel > 0 {
				log.Printf("Error encountered from Pingpp: %v\n", err)
			}
			return err
		}
	}

	if LogLevel > 2 {
		log.Printf("resBody from pingpp API: \n%v\n", string(resBody))
	}

	if v != nil {
		return JsonDecode(resBody, v)
	}

	return nil
}
