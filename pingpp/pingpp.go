package pingpp

import (
	"bytes"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	// 当前版本的api地址
	apiBase = "https://api.pingxx.com/v1"
	// 当前版本的api生成生成时间
	apiVersion = "2016-12-27"
	// httpclient等待时间
	defaultHTTPTimeout                  = 80 * time.Second
	TotalBackends                       = 1
	APIBackend         SupportedBackend = "api"
)

var (
	// 默认错误信息返回语言
	AcceptLanguage = "zh-CN"
	// ping++ api统一需要通过Authentication（http BasicAuth），需要在调用时赋值
	Key string

	// loglevel 是 debug 模式开关.
	// 0: no logging
	// 1: errors only
	// 2: errors + informational (default)
	// 3: errors + informational + debug
	LogLevel = 2

	//不用默认的defaultClient，自定义httpClient
	httpClient        = &http.Client{Timeout: defaultHTTPTimeout}
	backends          Backends
	AccountPrivateKey string
	OsInfo            string
)

type SupportedBackend string

// 定义统一后端处理接口
type Backend interface {
	Call(method, path, key string, body *url.Values, params []byte, v interface{}) error
}

// 获取当前sdk的版本
func Version() string {
	return "3.1.0"
}

/*2016-02-16 当前情况下没有代码调用了该函数
func SetHttpClient(client *http.Client) {
	httpClient = client
}*/

type Backends struct {
	API Backend
}

// 通过不同的参数获取不同的后端对象
func GetBackend(backend SupportedBackend) Backend {
	var ret Backend
	switch backend {
	case APIBackend:
		if backends.API == nil {
			backends.API = ApiBackend{backend, apiBase, httpClient}
		}

		ret = backends.API
	}
	return ret
}

//设定后端处理对象
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case APIBackend:
		backends.API = b
	}
}

func init() {
	var uname string
	switch runtime.GOOS {
	case "windows":
		uname = "windows"
	default:
		cmd := exec.Command("uname", "-a")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		cmd.Run()
		uname = out.String()
	}
	m := map[string]interface{}{
		"lang":             "golang",
		"lang_version":     runtime.Version(),
		"bindings_version": Version(),
		"publisher":        "pingpp",
		"uname":            uname,
	}
	content, _ := JsonEncode(m)
	OsInfo = string(content)
}
