package batchTransfer

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

func New(params *pingpp.BatchTransferParams) (*pingpp.BatchTransfer, error) {
	return getC().New(params)
}

func (c Client) New(params *pingpp.BatchTransferParams) (*pingpp.BatchTransfer, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	batchTransfer := &pingpp.BatchTransfer{}
	err := c.B.Call("POST", "/batch_transfers", c.Key, nil, paramsString, batchTransfer)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("New BatchTransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}

func Get(Id string) (*pingpp.BatchTransfer, error) {
	return getC().Get(Id)
}

func (c Client) Get(Id string) (*pingpp.BatchTransfer, error) {
	batchTransfer := &pingpp.BatchTransfer{}
	err := c.B.Call("GET", fmt.Sprintf("/batch_transfers/%s", Id), c.Key, nil, nil, batchTransfer)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Batchtransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}

func List(params *pingpp.PagingParams) (*pingpp.BatchTransferlList, error) {
	return getC().List(params)
}

func (c Client) List(params *pingpp.PagingParams) (*pingpp.BatchTransferlList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchTransferlList := &pingpp.BatchTransferlList{}
	err := c.B.Call("GET", "/batch_transfers", c.Key, body, nil, batchTransferlList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Batchtransfer List error: %v\n", err)
		}
	}
	return batchTransferlList, err
}
