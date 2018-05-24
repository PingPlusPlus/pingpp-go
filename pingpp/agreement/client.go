package agreement

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

// Client 请求
type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// New 创建签约
// @param appId string
// @param params AgreementParams
// @return Agreement
func New(params *pingpp.AgreementParams) (*pingpp.Agreement, error) {
	return getC().New(params)
}

// New 创建签约
func (c Client) New(params *pingpp.AgreementParams) (*pingpp.Agreement, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("AgreementParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create agreement is :\n %v\n ", string(paramsString))
	}

	agreement := &pingpp.Agreement{}
	err := c.B.Call("POST", "/agreements", c.Key, nil, paramsString, agreement)
	return agreement, err
}

// Get 查询签约对象
// @param agreementID 签约对象 ID
// @return Agreement
func Get(agreementID string) (*pingpp.Agreement, error) {
	return getC().Get(agreementID)
}

// Get 查询签约对象
func (c Client) Get(agreementID string) (*pingpp.Agreement, error) {
	agreement := &pingpp.Agreement{}
	err := c.B.Call("GET", fmt.Sprintf("/agreements/%s", agreementID), c.Key, nil, nil, agreement)
	return agreement, err
}

// List 查询签约对象列表
// @param app string
// @param status string
// @param params PagingParams
// @return AgreementList
func List(app, status string, params *pingpp.PagingParams) ([]pingpp.Agreement, error) {
	return getC().List(app, status, params)
}

// List 查询签约对象列表
func (c Client) List(app, status string, params *pingpp.PagingParams) ([]pingpp.Agreement, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	body.Add("app", app)
	if status != "" && status != "*" {
		body.Add("status", status)
	}
	agreements := []pingpp.Agreement{}
	err := c.B.Call("GET", "/agreements", c.Key, body, nil, &agreements)
	return agreements, err
}

// Update 更新签约商户对象
// @param agreementID string
// @param params AgreementUpdateParams
// @return Agreement
func Update(agreementID string, params *pingpp.AgreementUpdateParams) (*pingpp.Agreement, error) {
	return getC().Update(agreementID, params)
}

// Update 更新子商户对象
func (c Client) Update(agreementID string, params *pingpp.AgreementUpdateParams) (*pingpp.Agreement, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update Agreement  to pingpp is :\n %v\n ", string(paramsString))
	}

	agreement := &pingpp.Agreement{}
	err := c.B.Call("PUT", fmt.Sprintf("/agreements/%s", agreementID), c.Key, nil, paramsString, agreement)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return agreement, err
}
