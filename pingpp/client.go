package pingpp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultURL = "https://api.pingxx.com/v1/"

const apiversion = "2015-04-01"

const clientversion = "2.0.1"

type Backend interface {
	Call(method, path, key string, body *url.Values, v interface{}) ([]byte, error)
	CallJson(method, path, key string, body []byte, v interface{}) ([]byte, error)
}

type InternalBackend struct {
	url        string
	httpClient *http.Client
}

func newInternalBackend(url string, httpClient *http.Client) *InternalBackend {
	if url == "" {
		url = defaultURL
	}
	return &InternalBackend{
		url:        url,
		httpClient: httpClient,
	}
}

var Key string
var debug bool
var backend Backend

func setDebug(value bool) {
	debug = value
}

func getBackend() Backend {
	if backend == nil {
		backend = newInternalBackend("", http.DefaultClient)
	}
	return backend
}

func setBackend(b Backend) {
	backend = b
}

func getRequest(method string, path string, body *url.Values) (*http.Request, error) {
	if method == "GET" {
		if body != nil && len(*body) > 0 {
			path += "?" + body.Encode()
		}
		log.Printf("path : %v \n", path)
		return http.NewRequest(method, path, nil)
	} else if method == "POST" {
		if body != nil && len(*body) > 0 {
			return http.NewRequest("POST", path, bytes.NewBufferString(body.Encode()))
		}
	}
	return http.NewRequest(method, path, nil)
}

func (b *InternalBackend) CallJson(method string, path string, key string, body []byte, v interface{}) ([]byte, error) {
	if strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = b.url + path

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	if err != nil {
		log.Printf("Cannot create PingPP requests: %v\n", err)
		return nil, err
	}
	req.SetBasicAuth(key, "")
	req.Header.Add("Pingpp-Version", apiversion)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "PingPP v"+clientversion)
	req.Header.Add("Content-Length", strconv.Itoa(len(string(body))))

	log.Printf("Requesting %v %q\n", method, path)
	start := time.Now()
	res, err := b.httpClient.Do(req)
	if debug {
		log.Printf("Request completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to PingPP Failed %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	resultbytes := buf.Bytes()
	//log.Printf("Body: %v\n", string(resultbytes))
	if err != nil {
		log.Printf("Parse PingPP Response Failed %v\n", err)
		return nil, err
	}
	if res.StatusCode >= 400 {
		var errMap map[string]interface{}
		json.Unmarshal(resultbytes, &errMap)
		if e, found := errMap["error"]; !found {
			err := errors.New(string(resultbytes))
			log.Printf("Unparsable error returned from PingPP: %v\n", err)
			return nil, err
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

			log.Printf("Error encountered from PingPP: %v\n", err)
			return nil, err
		}
	}

	if debug {
		log.Printf("PingPP response: %q\n", resultbytes)
	}
	return resultbytes, err
}

func (b *InternalBackend) Call(method string, path string, key string, body *url.Values, v interface{}) ([]byte, error) {
	if strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = b.url + path

	req, err := getRequest(method, path, body)
	if res, ok := v.([]byte); ok {
		log.Printf("Request Body:%s\n", string(res))
		req.Body = ioutil.NopCloser(strings.NewReader(string(res)))
		req.Header.Add("Content-Type", "application/json")
	} else {
		log.Printf("Request Body:%s\n", body.Encode())
		req.Body = ioutil.NopCloser(strings.NewReader(body.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if err != nil {
		log.Printf("Cannot create PingPP requests: %v\n", err)
		return nil, err
	}
	req.SetBasicAuth(key, "")
	req.Header.Add("Pingpp-Version", apiversion)

	req.Header.Add("User-Agent", "PingPP v"+clientversion)
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	log.Printf("Requesting %v %q\n", method, path)
	start := time.Now()
	res, err := b.httpClient.Do(req)
	if debug {
		log.Printf("Request completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to PingPP Failed %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	resultbytes := buf.Bytes()
	//log.Printf("Body: %v\n", string(resultbytes))
	if err != nil {
		log.Printf("Parse PingPP Response Failed %v\n", err)
		return nil, err
	}
	if res.StatusCode >= 400 {
		var errMap map[string]interface{}
		json.Unmarshal(resultbytes, &errMap)
		if e, found := errMap["error"]; !found {
			err := errors.New(string(resultbytes))
			log.Printf("Unparsable error returned from PingPP: %v\n", err)
			return nil, err
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

			log.Printf("Error encountered from PingPP: %v\n", err)
			return nil, err
		}
	}

	if debug {
		log.Printf("PingPP response: %q\n", resultbytes)
	}
	return resultbytes, err
}
