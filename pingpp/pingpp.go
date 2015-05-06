package pingpp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const apiBase = "https://api.pingxx.com/v1"
const apiVersion = "2015-05-06"
const clientVersion = "2.1.0"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
// This is chosen to be consistent with the other pingpp language libraries and
// to coordinate with other timeouts configured in the pingpp infrastructure.
const defaultHTTPTimeout = 80 * time.Second

const TotalBackends = 1

type Backend interface {
	Call(method, path, key string, body *url.Values, params []byte, v interface{}) error
}

type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

type SupportedBackend string

const APIBackend SupportedBackend = "api"

// Backends are the currently supported endpoints.
type Backends struct {
	API Backend
}

// Key is the pingp API key used globally in the binding.
var Key string

// LogLevel is the logging level for this library.
// 0: no logging
// 1: errors only
// 2: errors + informational (default)
// 3: errors + informational + debug
var LogLevel = 2

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}
var backends Backends

func SetHttpClient(client *http.Client) {
	httpClient = client
}

func GetBackend(backend SupportedBackend) Backend {
	var ret Backend
	switch backend {
	case APIBackend:
		if backends.API == nil {
			backends.API = BackendConfiguration{backend, apiBase, httpClient}
		}

		ret = backends.API
	}
	return ret
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case APIBackend:
		backends.API = b
	}
}

func (s BackendConfiguration) Call(method, path, key string, form *url.Values, params []byte, v interface{}) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBuffer(params)
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

func (s *BackendConfiguration) NewRequest(method, path, key, contentType string, body io.Reader, params []byte) (*http.Request, error) {
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

	req.SetBasicAuth(key, "")
	req.Header.Add("Pingpp-version", apiVersion)
	req.Header.Add("User-Agent", "pingpp/v1 GoBindings/"+clientVersion)
	req.Header.Add("content-type", contentType)

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}
	start := time.Now()
	res, err := s.HTTPClient.Do(req)
	if LogLevel > 2 {
		log.Printf("Completed in %v\n", time.Since(start))
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
		// for some odd reason, the Erro structure doesn't unmarshal
		// initially I thought it was because it's a struct inside of a struct
		// but even after trying that, it still didn't work
		// so unmarshalling to a map for now and parsing the results manually
		// but should investigate later
		var errMap map[string]interface{}
		json.Unmarshal(resBody, &errMap)

		if e, found := errMap["error"]; !found {
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
		return json.Unmarshal(resBody, v)

	}

	return nil
}
