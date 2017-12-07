package app

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

/*
* 创建子商户对象
* @param appId string
* @param params SubAppParams
* @return SubApp
 */
func New(appId string, params *pingpp.SubAppParams) (*pingpp.SubApp, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.SubAppParams) (*pingpp.SubApp, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("SubAppParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create sub_app is :\n %v\n ", string(paramsString))
	}

	subApp := &pingpp.SubApp{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/sub_apps", appId), c.Key, nil, paramsString, subApp)
	return subApp, err
}

/*
* 查询子商户对象
* @param appId string
* @param subAppId string
* @return SubApp
 */
func Get(appId, subAppId string) (*pingpp.SubApp, error) {
	return getC().Get(appId, subAppId)
}

func (c Client) Get(appId, subAppId string) (*pingpp.SubApp, error) {
	subApp := &pingpp.SubApp{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), c.Key, nil, nil, subApp)
	return subApp, err
}

/*
* 查询子商户对象列表
* @param appId string
* @param params PagingParams
* @return SubAppList
 */
func List(appId string, params *pingpp.PagingParams) (*pingpp.SubAppList, error) {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *pingpp.PagingParams) (*pingpp.SubAppList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	subList := &pingpp.SubAppList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/sub_apps", appId), c.Key, body, nil, subList)
	return subList, err
}

/*
* 更新子商户对象
* @param appId string
* @param subAppId string
* @param params SubAppUpdateParams
* @return SubApp
 */
func Update(appId, subAppId string, params pingpp.SubAppUpdateParams) (*pingpp.SubApp, error) {
	return getC().Update(appId, subAppId, params)
}

func (c Client) Update(appId, subAppId string, params pingpp.SubAppUpdateParams) (*pingpp.SubApp, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update SubApp  to pingpp is :\n %v\n ", string(paramsString))
	}

	subApp := &pingpp.SubApp{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), c.Key, nil, paramsString, subApp)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return subApp, err
}

/*
* 删除子商户对象
* @param appId string
* @param subAppId string
* @return DeleteResult
 */
func Delete(appId, subAppId string) (*pingpp.DeleteResult, error) {
	return getC().Delete(appId, subAppId)
}

func (c Client) Delete(appId, subAppId string) (*pingpp.DeleteResult, error) {
	result := &pingpp.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), c.Key, nil, nil, result)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}
