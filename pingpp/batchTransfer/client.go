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

/*
* 创建批量转账
* @param params BatchTransferParams
* @return BatchTransfer
 */
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

/*
* 查询批量转账
* @param Id string
* @return BatchTransfer
 */
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

/*
* 查询批量转账列表
* @param params PagingParams
* @return BatchTransferlList
 */
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

/*
* 取消批量转账
* @param batchTransferId string
* @return BatchTransfer
 */
func Cancel(batchTransferId string) (*pingpp.BatchTransfer, error) {
	return getC().Cancel(batchTransferId)
}
func (c Client) Cancel(batchTransferId string) (*pingpp.BatchTransfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := pingpp.JsonEncode(cancelParams)

	batchTransfer := &pingpp.BatchTransfer{}
	err := c.B.Call("PUT", fmt.Sprintf("/batch_transfers/%s", batchTransferId), c.Key, nil, paramsString, batchTransfer)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf(" BatchTransfer error: %v\n", err)
		}
	}
	return batchTransfer, err
}
