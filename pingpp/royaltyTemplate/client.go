package royaltyTemplate

import (
	"fmt"
	"log"
	"net/url"

	"github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// 创建分润模板
func New(params *pingpp.RoyaltyTmplParams) (*pingpp.RoyaltyTmpl, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.RoyaltyTmplParams) (*pingpp.RoyaltyTmpl, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of create royalty_template request to pingpp is :\n %v\n ", string(paramsString))
	}

	royaltyTemplate := &pingpp.RoyaltyTmpl{}

	err := c.B.Call("POST", "/royalty_templates", c.Key, nil, paramsString, royaltyTemplate)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTemplate, err

}

//查询指定的分润模板
func Get(royaltyTmplId string) (*pingpp.RoyaltyTmpl, error) {
	return getC().Get(royaltyTmplId)
}

func (c Client) Get(royaltyTmplId string) (*pingpp.RoyaltyTmpl, error) {
	var body *url.Values
	body = &url.Values{}
	royaltyTmpl := &pingpp.RoyaltyTmpl{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, body, nil, royaltyTmpl)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get royalty Template error: %v\n", err)
		}
	}
	return royaltyTmpl, err
}

//更新分润模板
func Update(royaltyTmplId string, params *pingpp.RoyaltyTmplUpdateParams) (*pingpp.RoyaltyTmpl, error) {
	return getC().Update(royaltyTmplId, params)
}

func (c Client) Update(royaltyTmplId string, params *pingpp.RoyaltyTmplUpdateParams) (*pingpp.RoyaltyTmpl, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update royalty template to pingpp is :\n %v\n ", string(paramsString))
	}

	royaltyTmpl := &pingpp.RoyaltyTmpl{}

	err := c.B.Call("PUT", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, nil, paramsString, royaltyTmpl)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTmpl, err
}

//删除分润模板

func Delete(royaltyTmplId string) (*pingpp.DeleteResult, error) {
	return getC().Delete(royaltyTmplId)
}

func (c Client) Delete(royaltyTmplId string) (*pingpp.DeleteResult, error) {
	result := &pingpp.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, nil, nil, result)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Delete Royalty Template error: %v\n", err)
		}
	}
	return result, err
}

//查询分润模板列表
func List(params *pingpp.PagingParams) (*pingpp.RoyaltyTmplList, error) {
	return getC().List(params)
}
func (c Client) List(params *pingpp.PagingParams) (*pingpp.RoyaltyTmplList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTmplList := &pingpp.RoyaltyTmplList{}
	err := c.B.Call("GET", "/royalty_templates", c.Key, body, nil, royaltyTmplList)
	return royaltyTmplList, err
}
