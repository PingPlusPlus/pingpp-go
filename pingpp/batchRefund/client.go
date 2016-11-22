package batchRefund

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

func New(params *pingpp.BatchRefundParams) (*pingpp.BatchRefund, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.BatchRefundParams) (*pingpp.BatchRefund, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	batchRefund := &pingpp.BatchRefund{}
	err := c.B.Call("POST", "/batch_refunds", c.Key, nil, paramsString, batchRefund)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("New BatchRefunds error: %v\n", err)
		}
	}
	return batchRefund, err
}

func Get(Id string) (*pingpp.BatchRefund, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*pingpp.BatchRefund, error) {
	batchRefund := &pingpp.BatchRefund{}
	err := c.B.Call("GET", fmt.Sprintf("/batch_refunds/%s", Id), c.Key, nil, nil, batchRefund)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get BatchRefunds error: %v\n", err)
		}
	}
	return batchRefund, err
}

func List(params *pingpp.PagingParams) (*pingpp.BatchRefundlList, error) {
	return getC().List(params)
}

func (c Client) List(params *pingpp.PagingParams) (*pingpp.BatchRefundlList, error) {
	var body *url.Values
	body = &url.Values{}
	batchRefundlList := &pingpp.BatchRefundlList{}

	err := c.B.Call("GET", "/batch_refunds", c.Key, body, nil, batchRefundlList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get BatchRefunds List error: %v\n", err)
		}
	}
	return batchRefundlList, err
}
